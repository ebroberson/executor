package runoncehandler_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry-incubator/runtime-schema/bbs/fakebbs"
	"github.com/cloudfoundry-incubator/runtime-schema/models"
	steno "github.com/cloudfoundry/gosteno"
	"github.com/vito/gordon/fake_gordon"

	"github.com/cloudfoundry-incubator/executor/actionrunner/fakeactionrunner"
	"github.com/cloudfoundry-incubator/executor/log_streamer_factory"
	. "github.com/cloudfoundry-incubator/executor/runoncehandler"
	"github.com/cloudfoundry-incubator/executor/taskregistry/faketaskregistry"
)

var _ = Describe("RunOnceHandler", func() {
	var (
		handler *RunOnceHandler

		bbs              *fakebbs.FakeExecutorBBS
		runOnce          models.RunOnce
		fakeTaskRegistry *faketaskregistry.FakeTaskRegistry
		gordon           *fake_gordon.FakeGordon
		actionRunner     *fakeactionrunner.FakeActionRunner
		stack            string
	)

	BeforeEach(func() {
		bbs = fakebbs.NewFakeExecutorBBS()
		gordon = fake_gordon.New()
		actionRunner = fakeactionrunner.New()
		fakeTaskRegistry = faketaskregistry.New()
		stack = "penguin"

		runOnce = models.RunOnce{
			Guid:  "totally-unique",
			Stack: "penguin",
			Actions: []models.ExecutorAction{
				{
					models.RunAction{
						Script: "sudo reboot",
					},
				},
			},
		}

		logStreamerFactory := &log_streamer_factory.LogStreamerFactory{}

		handler = New(
			bbs,
			gordon,
			fakeTaskRegistry,
			actionRunner,
			logStreamerFactory,
			stack,
			steno.NewLogger("test-logger"),
		)

	})

	Describe("Handling a RunOnce", func() {
		JustBeforeEach(func() {
			handler.RunOnce(runOnce, "executor-id")
		})

		Context("when the RunOnce stack is empty", func() {
			BeforeEach(func() {
				runOnce.Stack = ""
			})

			It("should pick up the RunOnce", func() {
				Ω(fakeTaskRegistry.RegisteredRunOnces).Should(ContainElement(runOnce))
			})
		})

		Context("when the RunOnce stack does not match the executor's stack", func() {
			BeforeEach(func() {
				runOnce.Stack = "lion"
			})

			It("should not pick up the RunOnce", func() {
				Ω(fakeTaskRegistry.RegisteredRunOnces).ShouldNot(ContainElement(runOnce))
			})
		})
	})
})
