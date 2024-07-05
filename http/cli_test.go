package http

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"
)


var (
	dummyBlindAlerter = &SpyBlindAlerter{}
 	dummyPlayerStore = &StubPlayerStore{}
 	dummyStdIn = &bytes.Buffer{}
 	dummyStdOut = &bytes.Buffer{}
	dummyGame = &GameSpy{} 
)
const BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number"


// func TestCLI(t *testing.T) {

// 	t.Run("start game with 3 players and finish game with 'Chris' as winner", func(t *testing.T) {
// 		game := &GameSpy{}
// 		stdout := &bytes.Buffer{}

// 		in := userSends("3", "Chris wins")
// 		cli := NewCLI(in, stdout, game)

// 		cli.PlayPoker()

// 		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
// 		assertGameStartedWith(t, game, 3)
// 		assertFinishCalledWith(t, game, "Chris")
// 	})

// 	t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T) {
// 		game := &GameSpy{}

// 		in := userSends("8", "Cleo wins")
// 		cli := NewCLI(in, dummyStdOut, game)

// 		cli.PlayPoker()

// 		assertGameStartedWith(t, game, 8)
// 		assertFinishCalledWith(t, game, "Cleo")
// 	})

// 	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
// 		game := &GameSpy{}

// 		stdout := &bytes.Buffer{}
// 		in := userSends("pies")

// 		cli := NewCLI(in, stdout, game)
// 		cli.PlayPoker()

// 		assertGameNotStarted(t, game)
// 		assertMessagesSentToUser(t, stdout, PlayerPrompt, BadPlayerInputErrMsg)
// 	})
// }

func TestGame_Start(t *testing.T) {
	checkSchedulingCases := func(cases []ScheduledAlert, t *testing.T, blindAlerter *SpyBlindAlerter){
		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
	
				if len(blindAlerter.Alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
				}
	
				got := blindAlerter.Alerts[i]
				assertScheduledAlert(t, got, want)
			})
		}
	}
	t.Run("schedules alerts on game start for 5 players", func(t *testing.T) {
		blindAlerter := &SpyBlindAlerter{}
		game := NewTexasHoldem(blindAlerter, dummyPlayerStore)

		game.Start(5, dummyStdOut)

		cases := []ScheduledAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 10 * time.Minute, Amount: 200},
			{At: 20 * time.Minute, Amount: 300},
			{At: 30 * time.Minute, Amount: 400},
			{At: 40 * time.Minute, Amount: 500},
			{At: 50 * time.Minute, Amount: 600},
			{At: 60 * time.Minute, Amount: 800},
			{At: 70 * time.Minute, Amount: 1000},
			{At: 80 * time.Minute, Amount: 2000},
			{At: 90 * time.Minute, Amount: 4000},
			{At: 100 * time.Minute, Amount: 8000},
		}

		checkSchedulingCases(cases, t, blindAlerter)
	})

	t.Run("schedules alerts on game start for 7 players", func(t *testing.T) {
		blindAlerter := &SpyBlindAlerter{}
		game := NewTexasHoldem(blindAlerter, dummyPlayerStore)

		game.Start(7, dummyStdOut)


		cases := []ScheduledAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 12 * time.Minute, Amount: 200},
			{At: 24 * time.Minute, Amount: 300},
			{At: 36 * time.Minute, Amount: 400},
		}

		checkSchedulingCases(cases, t, blindAlerter)
	})
	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &GameSpy{}

		cli := NewCLI(in, stdout, game)
		cli.PlayPoker()

		gotPrompt := stdout.String()
		wantPrompt := PlayerPrompt

		if gotPrompt != wantPrompt {
			t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
		}

		if game.StartedWith != 7 {
			t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
		}
	})
	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("Pies\n")
		game := &GameSpy{}
	
		cli := NewCLI(in, stdout, game)
		cli.PlayPoker()
	
		if game.StartCalled {
			t.Errorf("game should not have started")
		}
	})
}

func assertScheduledAlert(t testing.TB, got, want ScheduledAlert) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}