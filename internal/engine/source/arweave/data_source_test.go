package arweave_test

// var initializeOnce sync.Once
//
// func initialize(t *testing.T) {
//	initializeOnce.Do(func() {
//		zap.ReplaceGlobals(zaptest.NewLogger(t))
//	})
// }

// func TestSource(t *testing.T) {
//	t.Parallel()
//
//	initialize(t)
//
//	type arguments struct {
//		conf *config.Module
//	}
//
//	var testcases []struct {
//		name      string
//		arguments arguments
//		want      require.ValueAssertionFunc
//		wantError require.ErrorAssertionFunc
//	}
//
//	for _, testcase := range testcases {
//		testcase := testcase
//
//		t.Run(testcase.name, func(t *testing.T) {
//			t.Parallel()
//
//			instance, err := arweave.NewSource(testcase.arguments.conf, nil, nil)
//			require.NoError(t, err, "new arweave dataSource")
//
//			var (
//				tasksChan = make(chan *engine.Tasks)
//				errorChan = make(chan error)
//			)
//
//			instance.Start(context.Background(), tasksChan, errorChan)
//
//			for {
//				select {
//				case tasks := <-tasksChan:
//					for _, task := range tasks.Tasks {
//						t.Logf("Task %s", task.ID())
//					}
//				case err := <-errorChan:
//					require.NoError(t, err)
//
//					return
//				}
//			}
//		})
//	}
// }
