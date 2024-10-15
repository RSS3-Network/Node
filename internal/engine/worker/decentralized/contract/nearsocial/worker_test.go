package nearsocial_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/rss3-network/node/config"
	source "github.com/rss3-network/node/internal/engine/source/near"
	worker "github.com/rss3-network/node/internal/engine/worker/decentralized/contract/nearsocial"
	"github.com/rss3-network/node/provider/near"
	workerx "github.com/rss3-network/node/schema/worker/decentralized"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestWorker_Near(t *testing.T) {
	t.Parallel()

	type arguments struct {
		task   *source.Task
		config *config.Module
	}

	testcases := []struct {
		name      string
		arguments arguments
		want      *activityx.Activity
		wantError require.ErrorAssertionFunc
	}{
		{
			name: "Social Post",
			arguments: arguments{
				task: &source.Task{
					Network: network.Near,
					Block: near.Block{
						Header: near.BlockHeader{
							Height:    130355240,
							GasPrice:  "100000000",
							Timestamp: 1728946260124390382,
						},
					},
					Transaction: near.Transaction{
						Transaction: near.TransactionDetails{
							SignerID:   "progr.near",
							ReceiverID: "social.near",
							Actions: []near.Action{
								{
									FunctionCall: &near.FunctionCallAction{
										Deposit:    "18160000000000000000000",
										MethodName: "set",
										Gas:        100000000000000,
										Args:       "eyJkYXRhIjp7InByb2dyLm5lYXIiOnsiaW5kZXgiOnsiaGFzaHRhZyI6Ilt7XCJrZXlcIjpcImJpdGNvaW5cIixcInZhbHVlXCI6e1widHlwZVwiOlwic29jaWFsXCIsXCJwYXRoXCI6XCJwcm9nci5uZWFyL3Bvc3QvbWFpblwifX0se1wia2V5XCI6XCJmdXR1cmVcIixcInZhbHVlXCI6e1widHlwZVwiOlwic29jaWFsXCIsXCJwYXRoXCI6XCJwcm9nci5uZWFyL3Bvc3QvbWFpblwifX1dIiwicG9zdCI6IntcImtleVwiOlwibWFpblwiLFwidmFsdWVcIjp7XCJ0eXBlXCI6XCJtZFwifX0ifSwicG9zdCI6eyJtYWluIjoie1widHlwZVwiOlwibWRcIixcInRleHRcIjpcIlRoZSB1c2VyIGV4cGVyaWVuY2UgaXMgcmVhbGx5IHN0aWxsIGF3ZnVsLCBsYWdnaW5nLCBtdXN0IHNpZ24gdG8gdm90ZSwgcGF5IHRvIHBvc3QsLi4uIEJ1dCBpdCdzIGEgdmVyeSBnb29kIFByb29mIG9mIENvbmNlcHQhIFxcblRoZSBjb25jZXB0IG9mIGN1c3RvbSB3aWRnZXQgaXMgcHJldHR5IGNvb2wgYnV0IHJlYWxseSBub3QgaW50dWl0aXZlLiBBcyBmb3IgYSBkZXYgaXQncyBub3QgbmF0dXJhbCB0byBmaWd1cmUgb3V0IGhvdyBpdCB3b3Jrcy4gQnV0IGl0J3MgYSBnb29kIGlkZWEg8J+klFxcbiNCaXRjb2luICNGdXR1cmVcXG5cXG5UaGlzIHBvc3QgY29zdCBtZSAwLjAyIE5lYXIg8J+RiSBGZWVsIGZyZWUgdG8gZG9uYXRlIGF0IG15IHVzZXJuYW1lIHdoaWNoIGlzIG15IGFkZHJlc3MgKFllcywgdGhpcyBmZWF0dXJlIGlzIHByZXR0eSBjb29sIDspXCJ9In19fX0=",
									},
								},
							},
							Hash: "4SvjarLF5SQ2oDW7qswFY7J8bDGymcX3B4ASZwWriPC1",
						},
						TransactionOutcome: near.TransactionOutcome{
							Outcome: near.Outcome{
								GasBurnt: 5146221498020,
							},
						},
					},
				},
			},
			want: &activityx.Activity{
				ID:        "4SvjarLF5SQ2oDW7qswFY7J8bDGymcX3B4ASZwWriPC1",
				Network:   network.Near,
				Index:     0,
				From:      "progr.near",
				To:        "social.near",
				Type:      typex.SocialPost,
				Platform:  workerx.PlatformNearSocial.String(),
				Timestamp: 1728946260,
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("514622149802000000000")),
					Decimal: 24,
				},
				Calldata: &activityx.Calldata{
					ParsedFunction: "set",
				},
				Status: true,
				Actions: []*activityx.Action{
					{
						Type:     typex.SocialPost,
						Platform: workerx.PlatformNearSocial.String(),
						From:     "progr.near",
						To:       "social.near",
						Metadata: metadata.SocialPost{
							Handle:    "progr.near",
							Body:      "The user experience is really still awful, lagging, must sign to vote, pay to post,... But it's a very good Proof of Concept! \nThe concept of custom widget is pretty cool but really not intuitive. As for a dev it's not natural to figure out how it works. But it's a good idea 🤔\n#Bitcoin #Future\n\nThis post cost me 0.02 Near 👉 Feel free to donate at my username which is my address (Yes, this feature is pretty cool ;)",
							Tags:      []string{"bitcoin", "future"},
							Timestamp: 1728946260,
						},
					},
				},
			},
			wantError: require.NoError,
		},
		{
			name: "Social Comment",
			arguments: arguments{
				task: &source.Task{
					Network: network.Near,
					Block: near.Block{
						Header: near.BlockHeader{
							Height:    130332693,
							GasPrice:  "100000000",
							Timestamp: 1728921368378666419,
						},
					},
					Transaction: near.Transaction{
						Transaction: near.TransactionDetails{
							SignerID:   "yarotska.near",
							ReceiverID: "social.near",
							Actions: []near.Action{
								{
									FunctionCall: &near.FunctionCallAction{
										Deposit:    "0",
										MethodName: "set",
										Gas:        100000000000000,
										Args:       "eyJkYXRhIjp7Inlhcm90c2thLm5lYXIiOnsicG9zdCI6eyJjb21tZW50Ijoie1widHlwZVwiOlwibWRcIixcInRleHRcIjpcIkhpIEBwaXZvcnRleC5uZWFyIOKAkyBjb25ncmF0dWxhdGlvbnMhIFRoaXMgbWVzc2FnZSBjb25maXJtcyB5b3VyIGZ1bmRpbmcgcmVxdWVzdCBhcHByb3ZhbCBieSB0aGUgRXZlbnRzIENvbW1pdHRlZS4gV2UncmUgZXhjaXRlZCB0byBzcG9uc29yIHlvdXIgd29yayEgVGhpcyBhcHByb3ZhbCBmb2xsb3dzIG91ciByZXZpZXcgcHJvY2VzcyBpbnZvbHZpbmcgdmFyaW91cyBXb3JrIEdyb3VwcyBhbmQgRXZlbnRzIENvbW1pdHRlZSBtZW1iZXJzLiBQbGVhc2Ugbm90ZSB0aGF0IHRoZSBmdW5kaW5nIGRpc3RyaWJ1dGlvbiBpcyBjb250aW5nZW50IG9uIHN1Y2Nlc3NmdWxseSBwYXNzaW5nIG91ciBLWUMvQiBhbmQgcGFwZXJ3b3JrIHByb2Nlc3MuXFxuXFxuKipJTVBPUlRBTlQgTk9URSoqXFxuVG8gZW5zdXJlIHRpbWVseSBkaXN0cmlidXRpb24gb2YgZnVuZHMsIHBsZWFzZSBjb21wbGV0ZSB0aGUgZnVuZGluZyByZXF1ZXN0IG5leHQgc3RlcHMgYW5kIHJlc3BvbmQgdG8gYW55IHJlcXVlc3RzIHdpdGhpbiAzMCBkYXlzIG9mIHJlY2VpdmluZyB0aGlzIG1lc3NhZ2UuIEZhaWx1cmUgdG8gZG8gc28gbWF5IHJlc3VsdCBpbiBjYW5jZWxpbmcgeW91ciBhcHByb3ZhbCBhbmQgd2l0aGhvbGRpbmcgdGhlIHBheW1lbnQuXFxuXFxuXFxuSGVyZeKAmXMgd2hhdCB0byBleHBlY3Q6XFxuXFxuKipGdW5kaW5nIFN0ZXBzKipcXG5cXG4xLiAqKktZQy9LWUIgVmVyaWZpY2F0aW9uOioqIEFuIEV2ZW50cyBDb21taXR0ZWUgbWVtYmVyIHdpbGwgbW92ZSB5b3VyIHByb3Bvc2FsIHRvIHRoZSBQYXltZW50IFByb2Nlc3NpbmcgU3RhZ2UgYW5kIHZlcmlmeSB0aGF0IHlvdSBoYXZlIGNvbXBsZXRlZCB2ZXJpZmljYXRpb24gdG8gZW5zdXJlIGNvbXBsaWFuY2UuIElmIHlvdSBhcmUgbm90IHZlcmlmaWVkLCB5b3VyIEV2ZW50cyBDb21taXR0ZWUgbWVtYmVyIHdpbGwgY29udGFjdCB5b3Ugb24gVGVsZWdyYW0gd2l0aCBpbnN0cnVjdGlvbnMgb24gaG93IHRvIHByb2NlZWQuIFRvIHJlY2VpdmUgZnVuZGluZywgeW91IG11c3QgZ2V0IHZlcmlmaWVkIHRocm91Z2ggRnJhY3RhbCwgYSB0cnVzdGVkIHRoaXJkLXBhcnR5IGlkZW50aWZpY2F0aW9uIHZlcmlmaWNhdGlvbiBzb2x1dGlvbi4gWW91ciB2ZXJpZmljYXRpb24gYmFkZ2UgaXMgdmFsaWQgZm9yIDM2NSBkYXlzIGFuZCBuZWVkcyByZW5ld2FsIHVwb24gZXhwaXJhdGlvbiBPUiBpZiB5b3VyIHBlcnNvbmFsIGluZm9ybWF0aW9uIGNoYW5nZXMsIHN1Y2ggYXMgeW91ciBuYW1lLCBhZGRyZXNzLCBvciBJRCBleHBpcmF0aW9uLlxcbjIuICoqSW5mb3JtYXRpb24gQ29sbGVjdGlvbjoqKiBPbmNlIHZlcmlmaWVkLCBhbiBFdmVudHMgQ29tbWl0dGVlIG1lbWJlciB3aWxsIGNvbnRhY3QgeW91IHZpYSBUZWxlZ3JhbSBhbmQgcmVxdWVzdCB0aGF0IHlvdSBjb21wbGV0ZSB0aGUgRnVuZGluZyBSZXF1ZXN0IEZvcm0gdXNpbmcgQWlydGFibGUuXFxuMy4gKipQcm9jZXNzaW5nOioqIE91ciBsZWdhbCB0ZWFtIHdpbGwgdmVyaWZ5IHlvdXIgYXBwbGljYXRpb24gZGV0YWlscyB0byBlbnN1cmUgY29tcGxpYW5jZS4gVGhleSB3aWxsIHRoZW4gc2VuZCB5b3UgYW4gZW1haWwgcmVxdWVzdGluZyB5b3VyIHNpZ25hdHVyZSBmb3IgdGhlIHVuZGVybHlpbmcgYWdyZWVtZW50IHZpYSBJcm9uY2xhZC5cXG40LiAqKkludm9pY2luZyAmIFBheW1lbnQ6KiogT25jZSB3ZSByZWNlaXZlIHlvdXIgc2lnbmVkIGFncmVlbWVudCwgb3VyIGZpbmFuY2UgdGVhbSB3aWxsIGVtYWlsIHlvdSBpbnN0cnVjdGlvbnMgdG8gc3VibWl0IHRoZSBmaW5hbCBpbnZvaWNlIHVzaW5nIFJlcXVlc3QgRmluYW5jZS4gT25jZSB3ZSByZWNlaXZlIHlvdXIgaW52b2ljZSwgb3VyIGZpbmFuY2UgdGVhbSB3aWxsIHNlbmQgYSB0ZXN0IHRyYW5zYWN0aW9uIGNvbmZpcm1hdGlvbiBlbWFpbC4gT25jZSB5b3UgY29uZmlybSB0aGUgdGVzdCB0cmFuc2FjdGlvbiwgd2Ugd2lsbCBkaXN0cmlidXRlIHRoZSBmdW5kcyBhbmQgcG9zdCBhIHBheW1lbnQgbGluayBvbiB5b3VyIHByb3Bvc2FsLlxcblxcbioqRnVuZGluZyBDb252ZXJzaW9uIE5vdGljZSoqXFxuXFxuT25jZSB5b3UgcmVjZWl2ZSB5b3VyIGZ1bmRpbmcsIHdlIHVyZ2UgeW91IHRvIGV4ZXJjaXNlIGNhdXRpb24gaWYgYXR0ZW1wdGluZyB0byBjb252ZXJ0IHlvdXIgZnVuZHMuIFNvbWUgdGhpcmQtcGFydHkgdG9vbHMgbWF5IGltcG9zZSBzaWduaWZpY2FudCBzd2FwcGluZyBmZWVzLlxcblxcbioqVmlzaWJpbGl0eSoqXFxuXFxuV2UgdHJhY2sgdGhlIGZ1bmRpbmcgcHJvY2VzcyBvbiBlYWNoIHByb3Bvc2FsIHVzaW5nIHRoZSB0aW1lbGluZSBhbmQgY29tbWVudHMuIEhvd2V2ZXIsIHlvdSBhcmUgd2VsY29tZSB0byByZWFjaCBvdXQgdG8gdGhlIEV2ZW50cyBDb21taXR0ZWUgbWVtYmVyIHdpdGggYW55IHF1ZXN0aW9ucy4gXFxuXFxuKipUaW1lbGluZSoqXFxuXFxuVHlwaWNhbGx5LCBmdW5kcyBhcmUgZGlzYnVyc2VkIHdpdGhpbiAxMCBidXNpbmVzcyBkYXlzLCBidXQgdGhlIHRpbWVsaW5lIGNhbiB2YXJ5IGRlcGVuZGluZyBvbiB0aGUgcHJvamVjdCdzIGNvbXBsZXhpdHkgYW5kIHBhcGVyd29yay4gWW91ciBFdmVudHMgQ29tbWl0dGVlIG1lbWJlciB3aWxsIGtlZXAgeW91IHVwZGF0ZWQuXFxuXCIsXCJpdGVtXCI6e1widHlwZVwiOlwic29jaWFsXCIsXCJwYXRoXCI6XCJldmVudHMtY29tbWl0dGVlLm5lYXIvcG9zdC9tYWluXCIsXCJibG9ja0hlaWdodFwiOjEzMDMxNjMyMH19In0sImluZGV4Ijp7ImNvbW1lbnQiOiJ7XCJrZXlcIjp7XCJ0eXBlXCI6XCJzb2NpYWxcIixcInBhdGhcIjpcImV2ZW50cy1jb21taXR0ZWUubmVhci9wb3N0L21haW5cIixcImJsb2NrSGVpZ2h0XCI6MTMwMzE2MzIwfSxcInZhbHVlXCI6e1widHlwZVwiOlwibWRcIn19Iiwibm90aWZ5IjoiW3tcImtleVwiOlwicGl2b3J0ZXgubmVhclwiLFwidmFsdWVcIjp7XCJ0eXBlXCI6XCJtZW50aW9uXCIsXCJpdGVtXCI6e1widHlwZVwiOlwic29jaWFsXCIsXCJwYXRoXCI6XCJ5YXJvdHNrYS5uZWFyL3Bvc3QvY29tbWVudFwifX19LHtcImtleVwiOlwicGl2b3J0ZXgubmVhclwiLFwidmFsdWVcIjp7XCJ0eXBlXCI6XCJwcm9wb3NhbC9yZXBseVwiLFwiaXRlbVwiOntcInR5cGVcIjpcInNvY2lhbFwiLFwicGF0aFwiOlwiZXZlbnRzLWNvbW1pdHRlZS5uZWFyL3Bvc3QvbWFpblwiLFwiYmxvY2tIZWlnaHRcIjoxMzAzMTYzMjB9LFwicHJvcG9zYWxcIjo4NSxcIndpZGdldEFjY291bnRJZFwiOlwiZXZlbnRzLWNvbW1pdHRlZS5uZWFyXCJ9fV0ifX19fQ==",
									},
								},
							},
							Hash: "BEbDBVrCw89nN6JSBBAu8MhBdfwvTiJMYao945cB4Qcg",
						},
						TransactionOutcome: near.TransactionOutcome{
							Outcome: near.Outcome{
								GasBurnt: 467752261535,
							},
						},
					},
				},
			},
			want: &activityx.Activity{
				ID:        "BEbDBVrCw89nN6JSBBAu8MhBdfwvTiJMYao945cB4Qcg",
				Network:   network.Near,
				Index:     0,
				From:      "yarotska.near",
				To:        "social.near",
				Type:      typex.SocialComment,
				Platform:  workerx.PlatformNearSocial.String(),
				Timestamp: 1728921368,
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("46775226153500000000")),
					Decimal: 24,
				},
				Calldata: &activityx.Calldata{
					ParsedFunction: "set",
				},
				Status: true,
				Actions: []*activityx.Action{
					{
						Type:     typex.SocialComment,
						Platform: workerx.PlatformNearSocial.String(),
						From:     "yarotska.near",
						To:       "social.near",
						Metadata: metadata.SocialPost{
							Handle:        "yarotska.near",
							Body:          "Hi @pivortex.near – congratulations! This message confirms your funding request approval by the Events Committee. We're excited to sponsor your work! This approval follows our review process involving various Work Groups and Events Committee members. Please note that the funding distribution is contingent on successfully passing our KYC/B and paperwork process.\n\n**IMPORTANT NOTE**\nTo ensure timely distribution of funds, please complete the funding request next steps and respond to any requests within 30 days of receiving this message. Failure to do so may result in canceling your approval and withholding the payment.\n\n\nHere’s what to expect:\n\n**Funding Steps**\n\n1. **KYC/KYB Verification:** An Events Committee member will move your proposal to the Payment Processing Stage and verify that you have completed verification to ensure compliance. If you are not verified, your Events Committee member will contact you on Telegram with instructions on how to proceed. To receive funding, you must get verified through Fractal, a trusted third-party identification verification solution. Your verification badge is valid for 365 days and needs renewal upon expiration OR if your personal information changes, such as your name, address, or ID expiration.\n2. **Information Collection:** Once verified, an Events Committee member will contact you via Telegram and request that you complete the Funding Request Form using Airtable.\n3. **Processing:** Our legal team will verify your application details to ensure compliance. They will then send you an email requesting your signature for the underlying agreement via Ironclad.\n4. **Invoicing & Payment:** Once we receive your signed agreement, our finance team will email you instructions to submit the final invoice using Request Finance. Once we receive your invoice, our finance team will send a test transaction confirmation email. Once you confirm the test transaction, we will distribute the funds and post a payment link on your proposal.\n\n**Funding Conversion Notice**\n\nOnce you receive your funding, we urge you to exercise caution if attempting to convert your funds. Some third-party tools may impose significant swapping fees.\n\n**Visibility**\n\nWe track the funding process on each proposal using the timeline and comments. However, you are welcome to reach out to the Events Committee member with any questions. \n\n**Timeline**\n\nTypically, funds are disbursed within 10 business days, but the timeline can vary depending on the project's complexity and paperwork. Your Events Committee member will keep you updated.\n",
							PublicationID: "events-committee.near/post/main-130316320",
							Timestamp:     1728921368,
							Target: &metadata.SocialPost{
								Handle:        "events-committee.near",
								PublicationID: "events-committee.near/post/main-130316320",
							},
						},
					},
				},
			},
			wantError: require.NoError,
		},
		{
			name: "Social Repost",
			arguments: arguments{
				task: &source.Task{
					Network: network.Near,
					Block: near.Block{
						Header: near.BlockHeader{
							Height:    130155103,
							GasPrice:  "100000000",
							Timestamp: 1728726130272950660,
						},
					},
					Transaction: near.Transaction{
						Transaction: near.TransactionDetails{
							SignerID:   "mbbevilacqua_caffepoesia.near",
							ReceiverID: "social.near",
							Actions: []near.Action{
								{
									FunctionCall: &near.FunctionCallAction{
										Deposit:    "0",
										MethodName: "set",
										Gas:        100000000000000,
										Args:       "eyJkYXRhIjp7Im1iYmV2aWxhY3F1YV9jYWZmZXBvZXNpYS5uZWFyIjp7ImluZGV4Ijp7InJlcG9zdCI6Ilt7XCJrZXlcIjpcIm1haW5cIixcInZhbHVlXCI6e1widHlwZVwiOlwicmVwb3N0XCIsXCJpdGVtXCI6e1widHlwZVwiOlwic29jaWFsXCIsXCJwYXRoXCI6XCJjYWZmZXBvZXNpYS5uZWFyL3Bvc3QvbWFpblwiLFwiYmxvY2tIZWlnaHRcIjoxMzAxNTQ5MjN9fX0se1wia2V5XCI6e1widHlwZVwiOlwic29jaWFsXCIsXCJwYXRoXCI6XCJjYWZmZXBvZXNpYS5uZWFyL3Bvc3QvbWFpblwiLFwiYmxvY2tIZWlnaHRcIjoxMzAxNTQ5MjN9LFwidmFsdWVcIjp7XCJ0eXBlXCI6XCJyZXBvc3RcIn19XSIsIm5vdGlmeSI6IntcImtleVwiOlwiY2FmZmVwb2VzaWEubmVhclwiLFwidmFsdWVcIjp7XCJ0eXBlXCI6XCJyZXBvc3RcIixcIml0ZW1cIjp7XCJ0eXBlXCI6XCJzb2NpYWxcIixcInBhdGhcIjpcImNhZmZlcG9lc2lhLm5lYXIvcG9zdC9tYWluXCIsXCJibG9ja0hlaWdodFwiOjEzMDE1NDkyM319fSJ9fX19",
									},
								},
							},
							Hash: "GKbNWy68JxQpHz6e5aM9HjMhEonD8UAYvcXuPM6HMtzg",
						},
						TransactionOutcome: near.TransactionOutcome{
							Outcome: near.Outcome{
								GasBurnt: 332807348085,
							},
						},
					},
				},
			},
			want: &activityx.Activity{
				ID:        "GKbNWy68JxQpHz6e5aM9HjMhEonD8UAYvcXuPM6HMtzg",
				Network:   network.Near,
				Index:     0,
				From:      "mbbevilacqua_caffepoesia.near",
				To:        "social.near",
				Type:      typex.SocialShare,
				Platform:  workerx.PlatformNearSocial.String(),
				Timestamp: 1728726130,
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("33280734808500000000")),
					Decimal: 24,
				},
				Calldata: &activityx.Calldata{
					ParsedFunction: "set",
				},
				Status: true,
				Actions: []*activityx.Action{
					{
						Type:     typex.SocialShare,
						Platform: workerx.PlatformNearSocial.String(),
						From:     "mbbevilacqua_caffepoesia.near",
						To:       "social.near",
						Metadata: metadata.SocialPost{
							Handle:    "mbbevilacqua_caffepoesia.near",
							Timestamp: 1728726130,
							Target: &metadata.SocialPost{
								Handle:        "caffepoesia.near",
								PublicationID: "caffepoesia.near/post/main-130154923",
							},
						},
					},
				},
			},
			wantError: require.NoError,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			instance, err := worker.NewWorker(testcase.arguments.config)
			require.NoError(t, err)

			feed, err := instance.Transform(ctx, testcase.arguments.task)
			testcase.wantError(t, err)

			t.Log(string(lo.Must(json.MarshalIndent(feed, "", "\x20\x20"))))

			require.Equal(t, testcase.want, feed)
		})
	}
}
