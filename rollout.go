package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"text/tabwriter"

	"github.com/coreos/updateservicectl/client/update/v1"
)

var (
	rolloutFlags struct {
		appId   StringFlag
		groupId StringFlag
		active  bool

		// linear rollouts
		frameSize int64
		duration  int64
	}

	cmdRollout = &Command{
		Name:    "rollout",
		Summary: "Operations on automated rollouts for a group.",
		Subcommands: []*Command{
			cmdRolloutCreate,
			cmdRolloutActivate,
			cmdRolloutDeactivate,
		},
		Run: rolloutGet,
	}
	cmdRolloutCreate = &Command{
		Name:    "rollout create",
		Usage:   "[OPTION]...",
		Summary: "Create an automated rollout.",
		Subcommands: []*Command{
			cmdRolloutLinear,
		},
	}
	cmdRolloutActivate = &Command{
		Name:    "rollout activate",
		Usage:   "[OPTION]...",
		Summary: "Activate a rollout for a group.",
		Run:     rolloutActivate,
	}
	cmdRolloutDeactivate = &Command{
		Name:    "rollout deactivate",
		Usage:   "[OPTION]...",
		Summary: "Deactivate a rollout for a group.",
		Run:     rolloutDeactivate,
	}

	// each type of rollout has it's own subcommand different arguments and
	// behavior, even though they all use the same API endpoint.
	cmdRolloutLinear = &Command{
		Name:    "rollout create linear",
		Usage:   "[OPTION]...",
		Summary: "Create a linear rollout.",
		Run:     rolloutLinear,
	}
)

func init() {
	// flags for getting a rollout
	cmdRollout.Flags.Var(&rolloutFlags.appId, "app-id",
		"Application containing the group the rollout is associated with.")
	cmdRollout.Flags.Var(&rolloutFlags.groupId, "group-id",
		"ID of the group the rollout is associated with.")

	// flags for activating a rollout
	cmdRolloutActivate.Flags.Var(&rolloutFlags.appId, "app-id",
		"Application containing the group the rollout is associated with.")
	cmdRolloutActivate.Flags.Var(&rolloutFlags.groupId, "group-id",
		"ID of the group the rollout is associated with.")

	// flags for deactivating a rollout
	cmdRolloutDeactivate.Flags.Var(&rolloutFlags.appId, "app-id",
		"Application containing the group the rollout is associated with.")
	cmdRolloutDeactivate.Flags.Var(&rolloutFlags.groupId, "group-id",
		"ID of the group the rollout is associated with.")

	// creating a linear rollout
	cmdRolloutLinear.Flags.Var(&rolloutFlags.appId, "app-id",
		"Application containing the group the rollout is associated with.")
	cmdRolloutLinear.Flags.Var(&rolloutFlags.groupId, "group-id",
		"ID of the group the rollout is associated with.")
	cmdRolloutLinear.Flags.Int64Var(&rolloutFlags.frameSize, "frame-size", 60,
		"Duration of a rollout step (or frame) in seconds (default 60 seconds)")
	cmdRolloutLinear.Flags.Int64Var(&rolloutFlags.duration, "duration", 86400,
		"Total duration for the rollout to go from 0% to 100%, in seconds (default 1 day)")
}

func displayRollout(out io.Writer, rollout *update.Rollout) {
	fmt.Fprintf(out, "rollout:\n")
	for i, frame := range rollout.Rollout {
		fmt.Fprintf(out, "Frame %d:\tPercent:\t%f\n", i, frame.Percent)
		fmt.Fprintf(out, "\tDuration:\t%d\n", frame.Duration)
	}
}

func rolloutGet(args []string, service *update.Service, out *tabwriter.Writer) int {
	if rolloutFlags.appId.Get() == nil ||
		rolloutFlags.groupId.Get() == nil {
		return ERROR_USAGE
	}

	call := service.Group.Rollout.Get(
		rolloutFlags.appId.String(), rolloutFlags.groupId.String())

	rollout, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}

	displayRollout(out, rollout)
	out.Flush()

	return OK
}

func setActive(service *update.Service, out *tabwriter.Writer, active bool) int {
	if rolloutFlags.appId.Get() == nil ||
		rolloutFlags.groupId.Get() == nil {
		return ERROR_USAGE
	}

	rolloutActive := &update.RolloutActive{
		Active: active,
	}

	setCall := service.Group.Rollout.Active.Set(
		rolloutFlags.appId.String(), rolloutFlags.groupId.String(), rolloutActive)

	rolloutActive, err := setCall.Do()
	if err != nil {
		log.Fatal(err)
	}

	if rolloutActive.Active {
		fmt.Fprintf(out, "rollout activated\n")
	} else {
		fmt.Fprintf(out, "rollout deactivated\n")
	}
	out.Flush()

	return OK
}

func rolloutActivate(args []string, service *update.Service, out *tabwriter.Writer) int {
	return setActive(service, out, true)
}

func rolloutDeactivate(args []string, service *update.Service, out *tabwriter.Writer) int {
	return setActive(service, out, false)
}

func generateLinear(appId, groupId string, frameSize, totalDuration int64) *update.Rollout {
	var frames []*update.Frame
	duration := totalDuration
	stepSize := 100 / math.Ceil(float64(duration)/float64(frameSize))
	percent := 0.0

	for {
		// we stop making frames when we are at the end of the rollout
		if duration <= 0 {
			break
		}

		// set our current frame and step size
		size := frameSize
		percent += stepSize
		if size >= duration {
			// if that's longer than the time we have left, this is the final
			// frame. use the rest of the duration to set us to 100%
			size = duration
		}

		// append our new frame to the end of the frames list
		frames = append(frames, &update.Frame{
			Duration: size,
			Percent:  percent,
		})

		// subtract our frame from the total duration
		duration -= size
	}

	// add one more frame to bring us up to 100%. the server will set this
	// frame, and the next time it checks, it will exit the rollout, since the
	// duration is set to 0.
	frames = append(frames, &update.Frame{
		Duration: 0,
		Percent:  100,
	})

	return &update.Rollout{
		AppId:   appId,
		GroupId: groupId,
		Rollout: frames,
	}
}

func rolloutLinear(args []string, service *update.Service, out *tabwriter.Writer) int {
	if rolloutFlags.appId.Get() == nil ||
		rolloutFlags.groupId.Get() == nil ||
		rolloutFlags.frameSize == 0 ||
		rolloutFlags.duration == 0 {
		return ERROR_USAGE
	}

	rollout := generateLinear(rolloutFlags.appId.String(), rolloutFlags.groupId.String(),
		rolloutFlags.frameSize, rolloutFlags.duration)

	call := service.Group.Rollout.Set(
		rolloutFlags.appId.String(), rolloutFlags.groupId.String(), rollout)

	rollout, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(out, "rollout set\n")
	out.Flush()

	return OK
}
