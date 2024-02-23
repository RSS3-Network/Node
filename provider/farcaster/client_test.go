package farcaster_test

//func TestClient_GetCastsByFid(t *testing.T) {
//	t.Parallel()
//
//	type arguments struct {
//		ctx       context.Context
//		fid       *int64
//		reverse   bool
//		pageSize  *int
//		pageToken string
//	}
//
//	testcases := []struct {
//		name      string
//		arguments arguments
//		want      farcaster.MessageResponse
//		wantError require.ErrorAssertionFunc
//	}{
//		{
//			name: "GetCastsByFid",
//			arguments: arguments{
//				ctx:       context.Background(),
//				fid:       lo.ToPtr(int64(14142)),
//				reverse:   true,
//				pageSize:  lo.ToPtr(1),
//				pageToken: "BQfP3SiXZvdfJ6bV8bW9Bbf5DM+hwxpX",
//			},
//			want: farcaster.MessageResponse{
//				Messages: []farcaster.Message{
//					{
//						Data: farcaster.MessageData{
//							Type:      farcaster.MessageTypeCastAdd.String(),
//							Fid:       14142,
//							Timestamp: 84397948,
//							Network:   "FARCASTER_NETWORK_MAINNET",
//							CastAddBody: &farcaster.CastAddBody{
//								EmbedsDeprecated:  []string{},
//								Mentions:          []uint64{},
//								Text:              "https://hoot.it/",
//								MentionsPositions: []int32{},
//								Embeds: []farcaster.Embed{
//									{
//										URL: "https://hoot.it/",
//									},
//								},
//							},
//						},
//						Hash:            "0x25c4ccb9af936d519e493754b3a866bbe17ed4b9",
//						HashScheme:      "HASH_SCHEME_BLAKE3",
//						Signature:       "bm2KGGzrmMPq33uDIvWOC7yTzW65MuDnlB9paQNs0uDLesfUsyR0Vm+DuXvNLpd9U8k2BMmW9nqFyIJ1cqMEBg==",
//						SignatureScheme: "SIGNATURE_SCHEME_ED25519",
//						Signer:          "0x4157bb4fe638e7143b2bf8f3b636d5322242f6586eac191fc78aa54675360195",
//					},
//				},
//				NextPageToken: "BQfPfCXEzLmvk21Rnkk3VLOoZrvhftS5",
//			},
//			wantError: require.NoError,
//		},
//	}
//
//	for _, testcase := range testcases {
//		testcase := testcase
//
//		t.Run(testcase.name, func(t *testing.T) {
//			t.Parallel()
//
//			farcasterClient, err := farcaster.NewClient(farcaster.EndpointMainnet)
//			testcase.wantError(t, err)
//
//			result, err := farcasterClient.GetCastsByFid(testcase.arguments.ctx, testcase.arguments.fid, testcase.arguments.reverse, testcase.arguments.pageSize, testcase.arguments.pageToken)
//			testcase.wantError(t, err)
//
//			data, err := json.MarshalIndent(result, "", "\x20\x20")
//			require.NoError(t, err)
//
//			t.Log(string(data))
//
//			require.Equal(t, testcase.want, *result)
//		})
//	}
//}
//
//func TestClient_GetCastByFidAndHash(t *testing.T) {
//	t.Parallel()
//
//	type arguments struct {
//		ctx  context.Context
//		fid  *int64
//		hash string
//	}
//
//	testcases := []struct {
//		name      string
//		arguments arguments
//		want      farcaster.Message
//		wantError require.ErrorAssertionFunc
//	}{
//		{
//			name: "GetCastByFidAndHash",
//			arguments: arguments{
//				ctx:  context.Background(),
//				fid:  lo.ToPtr(int64(14142)),
//				hash: "0x25c4ccb9af936d519e493754b3a866bbe17ed4b9",
//			},
//			want: farcaster.Message{
//				Data: farcaster.MessageData{
//					Type:      farcaster.MessageTypeCastAdd.String(),
//					Fid:       14142,
//					Timestamp: 84397948,
//					Network:   "FARCASTER_NETWORK_MAINNET",
//					CastAddBody: &farcaster.CastAddBody{
//						EmbedsDeprecated:  []string{},
//						Mentions:          []uint64{},
//						Text:              "https://hoot.it/",
//						MentionsPositions: []int32{},
//						Embeds: []farcaster.Embed{
//							{
//								URL: "https://hoot.it/",
//							},
//						},
//					},
//				},
//				Hash:            "0x25c4ccb9af936d519e493754b3a866bbe17ed4b9",
//				HashScheme:      "HASH_SCHEME_BLAKE3",
//				Signature:       "bm2KGGzrmMPq33uDIvWOC7yTzW65MuDnlB9paQNs0uDLesfUsyR0Vm+DuXvNLpd9U8k2BMmW9nqFyIJ1cqMEBg==",
//				SignatureScheme: "SIGNATURE_SCHEME_ED25519",
//				Signer:          "0x4157bb4fe638e7143b2bf8f3b636d5322242f6586eac191fc78aa54675360195",
//			},
//			wantError: require.NoError,
//		},
//	}
//
//	for _, testcase := range testcases {
//		testcase := testcase
//
//		t.Run(testcase.name, func(t *testing.T) {
//			t.Parallel()
//
//			farcasterClient, err := farcaster.NewClient(farcaster.EndpointMainnet)
//			testcase.wantError(t, err)
//
//			result, err := farcasterClient.GetCastByFidAndHash(testcase.arguments.ctx, testcase.arguments.fid, testcase.arguments.hash)
//			testcase.wantError(t, err)
//
//			data, err := json.MarshalIndent(result, "", "\x20\x20")
//			require.NoError(t, err)
//
//			t.Log(string(data))
//
//			require.Equal(t, testcase.want, *result)
//		})
//	}
//}
//
//func TestClient_GetVerificationsByFid(t *testing.T) {
//	t.Parallel()
//
//	type arguments struct {
//		ctx       context.Context
//		fid       *int64
//		pageToken string
//	}
//
//	testcases := []struct {
//		name      string
//		arguments arguments
//		want      farcaster.MessageResponse
//		wantError require.ErrorAssertionFunc
//	}{
//		{
//			name: "GetVerificationsByFid",
//			arguments: arguments{
//				ctx:       context.Background(),
//				fid:       lo.ToPtr(int64(14142)),
//				pageToken: "",
//			},
//			want: farcaster.MessageResponse{
//				Messages: []farcaster.Message{
//					{
//						Data: farcaster.MessageData{
//							Type:      farcaster.MessageTypeVerificationAddEthAddress.String(),
//							Fid:       14142,
//							Timestamp: 91165349,
//							Network:   "FARCASTER_NETWORK_MAINNET",
//							VerificationAddEthAddressBody: &farcaster.VerificationAddEthAddressBody{
//								Address:      "0x8888888198fbdc8c017870cc5d3c96d0cf15c4f0",
//								EthSignature: "c+kDPPR4mpcy0UZ9f5z5+kkMhpK+wPgS0VUgo7J3riAgMOYFJ48BhC5BMuAMU+Zsr03U9/lIrxgsCBr9E3J6RBw=",
//								BlockHash:    "0xea1823d0815b0e309b45b2fb317011165e72dec0f04a107381531cbb9a12091a",
//							},
//						},
//						Hash:            "0xf41b3e42c9c9d5356f9c2c80bf9fac6607b56ecd",
//						HashScheme:      "HASH_SCHEME_BLAKE3",
//						Signature:       "1gZjcAr85b2dXTAyA/uLRRPViuRnrB0iu5898XqrgLRSWVlxsUqxpMxQWTlIchY4DmiopsF4C4ET7K6beKl4AQ==",
//						SignatureScheme: "SIGNATURE_SCHEME_ED25519",
//						Signer:          "0x4157bb4fe638e7143b2bf8f3b636d5322242f6586eac191fc78aa54675360195",
//					},
//				},
//				NextPageToken: "",
//			},
//			wantError: require.NoError,
//		},
//	}
//
//	for _, testcase := range testcases {
//		testcase := testcase
//
//		t.Run(testcase.name, func(t *testing.T) {
//			t.Parallel()
//
//			farcasterClient, err := farcaster.NewClient(farcaster.EndpointMainnet)
//			testcase.wantError(t, err)
//
//			result, err := farcasterClient.GetVerificationsByFid(testcase.arguments.ctx, testcase.arguments.fid, testcase.arguments.pageToken)
//			testcase.wantError(t, err)
//
//			data, err := json.MarshalIndent(result, "", "\x20\x20")
//			require.NoError(t, err)
//
//			t.Log(string(data))
//
//			require.Equal(t, testcase.want, *result)
//		})
//	}
//}
//
//func TestClient_GetUserNameProofsByFid(t *testing.T) {
//	t.Parallel()
//
//	type arguments struct {
//		ctx context.Context
//		fid *int64
//	}
//
//	testcases := []struct {
//		name      string
//		arguments arguments
//		want      farcaster.ProofResponse
//		wantError require.ErrorAssertionFunc
//	}{
//		{
//			name: "GetUserNameProofsByFid",
//			arguments: arguments{
//				ctx: context.Background(),
//				fid: lo.ToPtr(int64(14142)),
//			},
//			want: farcaster.ProofResponse{
//				Proofs: []farcaster.UserNameProof{
//					{
//						Timestamp: 1687683557,
//						Name:      "brucexc",
//						Owner:     "0xe5d6216f0085a7f6b9b692e06cf5856e6fa41b55",
//						Signature: "kC0zcMH/icJF5dukVQC+I/MXIjHPAeLM7iEB7fHWFM5/tuU0ClTVFiy0eZoARxtvIga1IB9BsKK0jvQaLRAgyhs=",
//						Fid:       14142,
//						Type:      farcaster.UsernameTypeFname.String(),
//					},
//					{
//						Timestamp: 1693966153,
//						Name:      "brucexc.eth",
//						Owner:     "0x8888888198fbdc8c017870cc5d3c96d0cf15c4f0",
//						Signature: "DU3+mKv+xNkAF8xDeqn2gJK/qkN3Pv4S9crym279awN6edVMKkQeMUYyvlzXusijWDQPrYkuN4sj8hf2ruXTZhw=",
//						Fid:       14142,
//						Type:      farcaster.UsernameTypeEnsL1.String(),
//					},
//				},
//			},
//			wantError: require.NoError,
//		},
//	}
//
//	for _, testcase := range testcases {
//		testcase := testcase
//
//		t.Run(testcase.name, func(t *testing.T) {
//			t.Parallel()
//
//			farcasterClient, err := farcaster.NewClient(farcaster.EndpointMainnet)
//			testcase.wantError(t, err)
//
//			result, err := farcasterClient.GetUserNameProofsByFid(testcase.arguments.ctx, testcase.arguments.fid)
//			testcase.wantError(t, err)
//
//			data, err := json.MarshalIndent(result, "", "\x20\x20")
//			require.NoError(t, err)
//
//			t.Log(string(data))
//
//			require.Equal(t, testcase.want, *result)
//		})
//	}
//}
//
//func TestClient_GetUserDataByFid(t *testing.T) {
//	t.Parallel()
//
//	type arguments struct {
//		ctx       context.Context
//		fid       *int64
//		pageToken string
//	}
//
//	testcases := []struct {
//		name      string
//		arguments arguments
//		want      farcaster.MessageResponse
//		wantError require.ErrorAssertionFunc
//	}{
//		{
//			name: "GetUserDataByFid",
//			arguments: arguments{
//				ctx:       context.Background(),
//				fid:       lo.ToPtr(int64(14142)),
//				pageToken: "",
//			},
//			want: farcaster.MessageResponse{
//				Messages: []farcaster.Message{
//					{
//						Data: farcaster.MessageData{
//							Type:      farcaster.MessageTypeUserDataAdd.String(),
//							Fid:       14142,
//							Timestamp: 78224506,
//							Network:   "FARCASTER_NETWORK_MAINNET",
//							UserDataBody: &farcaster.UserDataBody{
//								Type:  farcaster.UserDataTypeDisplay.String(),
//								Value: "bruce",
//							},
//						},
//						Hash:            "0x3342fa7d3e9c7a18a0e671ad7c163ada2543b934",
//						HashScheme:      "HASH_SCHEME_BLAKE3",
//						Signature:       "OlWA6GWLieSwmlM0HyiTUcYfLJCNTThDKGfz6Ndham77tnAQA7rz5BPtyGF98btNVkG0AFhhSsWWQb2SFbVUDg==",
//						SignatureScheme: "SIGNATURE_SCHEME_ED25519",
//						Signer:          "0x4157bb4fe638e7143b2bf8f3b636d5322242f6586eac191fc78aa54675360195",
//					},
//					{
//						Data: farcaster.MessageData{
//							Type:      farcaster.MessageTypeUserDataAdd.String(),
//							Fid:       14142,
//							Timestamp: 78224506,
//							Network:   "FARCASTER_NETWORK_MAINNET",
//							UserDataBody: &farcaster.UserDataBody{
//								Type:  farcaster.UserDataTypeBio.String(),
//								Value: "index from rss3",
//							},
//						},
//						Hash:            "0x73d0c2fd7463b46a87108034430e39e14eb1663a",
//						HashScheme:      "HASH_SCHEME_BLAKE3",
//						Signature:       "3TMlQJnEHXq45AywNvAqDvl7PJcTeT/Go/EVC0mDY1OsmPTf2XN9MR98qoZGdv1TRp+VqNz32S89zEuu17ImAQ==",
//						SignatureScheme: "SIGNATURE_SCHEME_ED25519",
//						Signer:          "0x4157bb4fe638e7143b2bf8f3b636d5322242f6586eac191fc78aa54675360195",
//					},
//					{
//						Data: farcaster.MessageData{
//							Type:      farcaster.MessageTypeUserDataAdd.String(),
//							Fid:       14142,
//							Timestamp: 78224519,
//							Network:   "FARCASTER_NETWORK_MAINNET",
//							UserDataBody: &farcaster.UserDataBody{
//								Type:  farcaster.UserDataTypePfp.String(),
//								Value: "https://i.seadn.io/gcs/files/c23d1b181620bb924fef8706fe745a95.png?w=500&auto=format",
//							},
//						},
//						Hash:            "0x2c061ac83aad2f88df5f3b6a1e91f390d9fc2b5c",
//						HashScheme:      "HASH_SCHEME_BLAKE3",
//						Signature:       "yKHuYZCy/m7xwJzE1ONqyOekLlovEZHUOaJAotGrzegOeBQ82KQjihK80mxakIGwV55cVutdSIn+khmFTDEJAA==",
//						SignatureScheme: "SIGNATURE_SCHEME_ED25519",
//						Signer:          "0x4157bb4fe638e7143b2bf8f3b636d5322242f6586eac191fc78aa54675360195",
//					},
//					{
//						Data: farcaster.MessageData{
//							Type:      farcaster.MessageTypeUserDataAdd.String(),
//							Fid:       14142,
//							Timestamp: 84506961,
//							Network:   "FARCASTER_NETWORK_MAINNET",
//							UserDataBody: &farcaster.UserDataBody{
//								Type:  farcaster.UserDataTypeUsername.String(),
//								Value: "brucexc.eth",
//							},
//						},
//						Hash:            "0x457cde2bd8c633c13c75b1e71349bdfdabe07057",
//						HashScheme:      "HASH_SCHEME_BLAKE3",
//						Signature:       "+YNv2lpzidgP3DMOUgM+hnMVTiRQvww1iti5yYlGQcRya1cmd2v2hIzBPbRr58VuhGfg5PHlS9KI1NYdho+OBQ==",
//						SignatureScheme: "SIGNATURE_SCHEME_ED25519",
//						Signer:          "0x4157bb4fe638e7143b2bf8f3b636d5322242f6586eac191fc78aa54675360195",
//					},
//				},
//				NextPageToken: "",
//			},
//			wantError: require.NoError,
//		},
//	}
//
//	for _, testcase := range testcases {
//		testcase := testcase
//
//		t.Run(testcase.name, func(t *testing.T) {
//			t.Parallel()
//
//			farcasterClient, err := farcaster.NewClient(farcaster.EndpointMainnet)
//			testcase.wantError(t, err)
//
//			result, err := farcasterClient.GetUserDataByFid(testcase.arguments.ctx, testcase.arguments.fid, testcase.arguments.pageToken)
//			testcase.wantError(t, err)
//
//			data, err := json.MarshalIndent(result, "", "\x20\x20")
//			require.NoError(t, err)
//
//			t.Log(string(data))
//
//			require.Equal(t, testcase.want, *result)
//		})
//	}
//}
//
//func TestClient_GetUserDataByFidAndType(t *testing.T) {
//	t.Parallel()
//
//	type arguments struct {
//		ctx          context.Context
//		fid          *int64
//		userDataType string
//	}
//
//	testcases := []struct {
//		name      string
//		arguments arguments
//		want      farcaster.Message
//		wantError require.ErrorAssertionFunc
//	}{
//		{
//			name: "GetUserDataByFid",
//			arguments: arguments{
//				ctx:          context.Background(),
//				fid:          lo.ToPtr(int64(14142)),
//				userDataType: farcaster.UserDataTypeDisplay.String(),
//			},
//			want: farcaster.Message{
//				Data: farcaster.MessageData{
//					Type:      farcaster.MessageTypeUserDataAdd.String(),
//					Fid:       14142,
//					Timestamp: 78224506,
//					Network:   "FARCASTER_NETWORK_MAINNET",
//					UserDataBody: &farcaster.UserDataBody{
//						Type:  farcaster.UserDataTypeDisplay.String(),
//						Value: "bruce",
//					},
//				},
//				Hash:            "0x3342fa7d3e9c7a18a0e671ad7c163ada2543b934",
//				HashScheme:      "HASH_SCHEME_BLAKE3",
//				Signature:       "OlWA6GWLieSwmlM0HyiTUcYfLJCNTThDKGfz6Ndham77tnAQA7rz5BPtyGF98btNVkG0AFhhSsWWQb2SFbVUDg==",
//				SignatureScheme: "SIGNATURE_SCHEME_ED25519",
//				Signer:          "0x4157bb4fe638e7143b2bf8f3b636d5322242f6586eac191fc78aa54675360195",
//			},
//			wantError: require.NoError,
//		},
//	}
//
//	for _, testcase := range testcases {
//		testcase := testcase
//
//		t.Run(testcase.name, func(t *testing.T) {
//			t.Parallel()
//
//			farcasterClient, err := farcaster.NewClient(farcaster.EndpointMainnet)
//			testcase.wantError(t, err)
//
//			result, err := farcasterClient.GetUserDataByFidAndType(testcase.arguments.ctx, testcase.arguments.fid, testcase.arguments.userDataType)
//			testcase.wantError(t, err)
//
//			data, err := json.MarshalIndent(result, "", "\x20\x20")
//			require.NoError(t, err)
//
//			t.Log(string(data))
//
//			require.Equal(t, testcase.want, *result)
//		})
//	}
//}
//
//func TestClient_GetEvents(t *testing.T) {
//	t.Parallel()
//
//	type arguments struct {
//		ctx         context.Context
//		fromEventID *int64
//	}
//
//	testcases := []struct {
//		name      string
//		arguments arguments
//		want      bool
//		wantError require.ErrorAssertionFunc
//	}{
//		{
//			name: "GetEvents",
//			arguments: arguments{
//				ctx:         context.Background(),
//				fromEventID: nil,
//			},
//			wantError: require.NoError,
//		},
//	}
//
//	for _, testcase := range testcases {
//		testcase := testcase
//
//		t.Run(testcase.name, func(t *testing.T) {
//			t.Parallel()
//
//			farcasterClient, err := farcaster.NewClient(farcaster.EndpointMainnet)
//			testcase.wantError(t, err)
//
//			result, err := farcasterClient.GetEvents(testcase.arguments.ctx, testcase.arguments.fromEventID)
//			testcase.wantError(t, err)
//
//			t.Log(len(result.Events))
//
//			require.Greater(t, len(result.Events), 0)
//		})
//	}
//}
//
//func TestClient_GetFids(t *testing.T) {
//	t.Parallel()
//
//	type arguments struct {
//		ctx      context.Context
//		reverse  bool
//		pageSize *int
//	}
//
//	testcases := []struct {
//		name      string
//		arguments arguments
//		want      farcaster.FidResponse
//		wantError require.ErrorAssertionFunc
//	}{
//		{
//			name: "GetFids",
//			arguments: arguments{
//				ctx:      context.Background(),
//				reverse:  true,
//				pageSize: lo.ToPtr(1),
//			},
//			wantError: require.NoError,
//		},
//	}
//
//	for _, testcase := range testcases {
//		testcase := testcase
//
//		t.Run(testcase.name, func(t *testing.T) {
//			t.Parallel()
//
//			farcasterClient, err := farcaster.NewClient(farcaster.EndpointMainnet)
//			testcase.wantError(t, err)
//
//			result, err := farcasterClient.GetFids(testcase.arguments.ctx, testcase.arguments.reverse, testcase.arguments.pageSize)
//			testcase.wantError(t, err)
//
//			t.Log(result.Fids[0])
//
//			require.Greater(t, result.Fids[0], uint64(196331))
//		})
//	}
//}
//
//func TestClient_GetReactionsByFid(t *testing.T) {
//	t.Parallel()
//
//	type arguments struct {
//		ctx          context.Context
//		fid          *int64
//		reverse      bool
//		pageSize     *int
//		pageToken    string
//		reactionType string
//	}
//
//	testcases := []struct {
//		name      string
//		arguments arguments
//		want      farcaster.MessageResponse
//		wantError require.ErrorAssertionFunc
//	}{
//		{
//			name: "GetReactionsByFid",
//			arguments: arguments{
//				ctx:          context.Background(),
//				fid:          lo.ToPtr(int64(14142)),
//				reverse:      true,
//				pageSize:     lo.ToPtr(1),
//				pageToken:    "BWWo+w1isWEOfdF3NjzTVxv+mhEtYYXe",
//				reactionType: farcaster.ReactionTypeRecast.String(),
//			},
//			want: farcaster.MessageResponse{
//				Messages: []farcaster.Message{
//					{
//						Data: farcaster.MessageData{
//							Type:      farcaster.MessageTypeReactionAdd.String(),
//							Fid:       14142,
//							Timestamp: 90547939,
//							Network:   "FARCASTER_NETWORK_MAINNET",
//							ReactionBody: &farcaster.ReactionBody{
//								Type: farcaster.ReactionTypeRecast.String(),
//								TargetCastID: farcaster.CastID{
//									Fid:  23901,
//									Hash: "0x05536f622dbedaa75ae1bd8ec9ff98bb0df10894",
//								},
//							},
//						},
//						Hash:            "0x2931794842a5f3a152bff66cc010120cfe9c6102",
//						HashScheme:      "HASH_SCHEME_BLAKE3",
//						Signature:       "L5DbR5RqsowyJxxfhgIfdh711e3GyxT+pLCCOTYRNLQ1+UyKikiExlXGXVcWlFHxX2grFNIfFmKeTvsFx9lfAg==",
//						SignatureScheme: "SIGNATURE_SCHEME_ED25519",
//						Signer:          "0x4157bb4fe638e7143b2bf8f3b636d5322242f6586eac191fc78aa54675360195",
//					},
//				},
//				NextPageToken: "BWWm4ykxeUhCpfOhUr/2bMAQEgz+nGEC",
//			},
//			wantError: require.NoError,
//		},
//	}
//
//	for _, testcase := range testcases {
//		testcase := testcase
//
//		t.Run(testcase.name, func(t *testing.T) {
//			t.Parallel()
//
//			farcasterClient, err := farcaster.NewClient(farcaster.EndpointMainnet)
//			testcase.wantError(t, err)
//
//			result, err := farcasterClient.GetReactionsByFid(testcase.arguments.ctx, testcase.arguments.fid, testcase.arguments.reverse, testcase.arguments.pageSize, testcase.arguments.pageToken, testcase.arguments.reactionType)
//			testcase.wantError(t, err)
//
//			data, err := json.MarshalIndent(result, "", "\x20\x20")
//			require.NoError(t, err)
//
//			t.Log(string(data))
//
//			require.Equal(t, testcase.want, *result)
//		})
//	}
//}
//
//func TestClient_GetReaction(t *testing.T) {
//	t.Parallel()
//
//	type arguments struct {
//		ctx          context.Context
//		fid          *int64
//		targetFid    *int64
//		targetHash   string
//		reactionType string
//	}
//
//	testcases := []struct {
//		name      string
//		arguments arguments
//		want      farcaster.Message
//		wantError require.ErrorAssertionFunc
//	}{
//		{
//			name: "GetReaction",
//			arguments: arguments{
//				ctx:          context.Background(),
//				fid:          lo.ToPtr(int64(14142)),
//				targetFid:    lo.ToPtr(int64(23901)),
//				targetHash:   "0x05536f622dbedaa75ae1bd8ec9ff98bb0df10894",
//				reactionType: farcaster.ReactionTypeRecast.String(),
//			},
//			want: farcaster.Message{
//
//				Data: farcaster.MessageData{
//					Type:      farcaster.MessageTypeReactionAdd.String(),
//					Fid:       14142,
//					Timestamp: 90547939,
//					Network:   "FARCASTER_NETWORK_MAINNET",
//					ReactionBody: &farcaster.ReactionBody{
//						Type: farcaster.ReactionTypeRecast.String(),
//						TargetCastID: farcaster.CastID{
//							Fid:  23901,
//							Hash: "0x05536f622dbedaa75ae1bd8ec9ff98bb0df10894",
//						},
//					},
//				},
//				Hash:            "0x2931794842a5f3a152bff66cc010120cfe9c6102",
//				HashScheme:      "HASH_SCHEME_BLAKE3",
//				Signature:       "L5DbR5RqsowyJxxfhgIfdh711e3GyxT+pLCCOTYRNLQ1+UyKikiExlXGXVcWlFHxX2grFNIfFmKeTvsFx9lfAg==",
//				SignatureScheme: "SIGNATURE_SCHEME_ED25519",
//				Signer:          "0x4157bb4fe638e7143b2bf8f3b636d5322242f6586eac191fc78aa54675360195",
//			},
//			wantError: require.NoError,
//		},
//	}
//
//	for _, testcase := range testcases {
//		testcase := testcase
//
//		t.Run(testcase.name, func(t *testing.T) {
//			t.Parallel()
//
//			farcasterClient, err := farcaster.NewClient(farcaster.EndpointMainnet)
//			testcase.wantError(t, err)
//
//			result, err := farcasterClient.GetReaction(testcase.arguments.ctx, testcase.arguments.fid, testcase.arguments.targetFid, testcase.arguments.targetHash, testcase.arguments.reactionType)
//			testcase.wantError(t, err)
//
//			data, err := json.MarshalIndent(result, "", "\x20\x20")
//			require.NoError(t, err)
//
//			t.Log(string(data))
//
//			require.Equal(t, testcase.want, *result)
//		})
//	}
//}
//
//func TestClient_Func(t *testing.T) {
//	t.Parallel()
//
//	t.Run("covert Farcaster time to timestamp", func(t *testing.T) {
//		t.Parallel()
//
//		require.Equal(t, farcaster.CovertFarcasterTimeToTimestamp(78224681), int64(1687683881))
//	})
//}
