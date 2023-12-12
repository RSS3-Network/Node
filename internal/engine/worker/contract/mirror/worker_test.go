package mirror_test

import (
	"context"
	"testing"

	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	source "github.com/naturalselectionlabs/rss3-node/internal/engine/source/arweave"
	worker "github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/mirror"
	"github.com/naturalselectionlabs/rss3-node/provider/arweave"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/naturalselectionlabs/rss3-node/schema/metadata"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestWorker_Arweave(t *testing.T) {
	t.Parallel()

	type arguments struct {
		task   *source.Task
		config *engine.Config
	}

	testcases := []struct {
		name      string
		arguments arguments
		want      *schema.Feed
		wantError require.ErrorAssertionFunc
	}{
		{
			name: "Mirror Post",
			arguments: arguments{
				task: &source.Task{
					Chain: filter.ChainArweaveMainnet,
					Transaction: arweave.Transaction{
						ID:       "lW0AMDN2RgOeqULk-u6Tv0wfZWpx9MfkrmqQQU-Mvuo",
						Reward:   "27798031",
						Quantity: "0",
						Owner:    "x3EoIMrB1qOrAivgqa1fi3SwJQgZdVuLgZhnvB0NLVngTIfB9yytGQVr9jJ4w4zn6q80T_L8aU5yXVhB12cfoABwP6AYDbvLefXLCGFgUcEcwrQ_bhWmDpxEjU2sMu4qrlI2PVj_3y1-Cx1wXldVD3i2JT81BeXacj8f19ZpKFZHaSvLM7jHm5STqwNaZDhhNnwR9qwe5kKkjafvOVDF9ayp7NHLOsEq5s8rcaQKZMYR_yEbuSfE4jzuq5F3Cwcp18bzooOqgtpCi1LozAlWepYywrxuhc-TxxObar2FVZU7bHyuNMD_aLoZ7mYqHUDGwPocqMqZVa6zQPAm7J1mnyI-fab-3E9vlZ3ZEO-PD21EAo_ALidmC3jR35tzWTGx6y7D4W1J_biJc_iKBybmsOH5W0K31CMW1sO2P2FKndUVz8vzedYvNbjYJhFFqAaEcn2x4mj15n2d75_DyEqmyEMXZ6MAg2-hUixYky7Gk7hpokwbwWEu35FsEIh4VKoanxJPgZ7JCWzt7su5su891RVf0Aodvr8UR971o1xgGaTR5ChLBC4SjyswcRTzNIb5LiPoJK4Sp03wAQW7dGgbKo8W5XjLsI-KWqypEyRkQYy-GYzwIWZ-4_5qzARaZOkVS_mn2aZzbsxxZNWxtjpJVNOd8olr0Fo88nn1xnLu2b0",
						Data:     "eyJjb250ZW50Ijp7ImJvZHkiOiJTdGF0ZWZ1bCBXb3JrcyBpcyBleGNpdGVkIHRvIHB1YmxpY2x5IGludHJvZHVjZSBvdXIgZmlyc3QgcHJvamVjdCwgYSBsaW1pdGVkIHJ1biBwcmludCBib29rIGZvY3VzZWQgb24gdGhlIHBlb3BsZSBhbmQgdGVhbXMgdGhhdCBicm91Z2h0IHRoZSBFdGhlcmV1bSBCZWFjb24gQ2hhaW4gdG8gbGlmZS5cblxuIyMjIFdobyBpcyBTdGF0ZWZ1bCBXb3Jrcz9cblxuV2UgYXJlIGEgZ3JvdXAgb2YgcGFzc2lvbmF0ZSBsb25ndGltZSBjb21tdW5pdHkgbWVtYmVycyBpbnRlcmVzdGVkIGluIGV4cGxvcmluZyBhbmQgY3JlYXRpbmcgRXRoZXJldW0gY3VsdHVyZTogcmVhY2ggb3V0IHRvIFtAc3RhdGVmdWx3b3Jrc10oaHR0cHM6Ly90d2l0dGVyLmNvbS9TdGF0ZWZ1bFdvcmtzKSBvbiBUZWxlZ3JhbSBvciBUd2l0dGVyIHdpdGggcXVlc3Rpb25zLlxuXG4jIyMgQm9vayBEZXRhaWxzXG5cblRoaXMgYm9vayB3aWxsIGZlYXR1cmUgcHJvZmlsZXMsIHJlZmxlY3Rpb25zLCBhbmQgc3RvcmllcyBmcm9tIEV0aDIgaW1wbGVtZW50ZXJzIGFuZCByZXNlYXJjaGVycy4gV2Ugd2FudCB0byBzaGFyZSB0aGUgc3RvcmllcyBvZiB0aGUgcGVvcGxlIHRoYXQgbWFkZSB0aGlzIHBvc3NpYmxlLiBXZSBoYXZlIGJlZW4gd29ya2luZyBvbiB0aGlzIGZvciB0aGUgcGFzdCBmZXcgbW9udGhzLCBhbmQgaGF2ZSBtYWRlIHNpZ25pZmljYW50IHByb2dyZXNzIHRvd2FyZHMgZGVmaW5pbmcgYW5kIHJlYWxpemluZyB0aGUgdmlzaW9uLlxuXG5TbyBmYXIsIHdlIGhhdmUgZ2F0aGVyZWQgcHJvZmlsZXMgZnJvbSBvdmVyICoqMjUgRXRoMiBjb250cmlidXRvcnMqKiwgd2l0aCBtYW55IG1vcmUgdG8gY29tZS4gV2UncmUgYW50aWNpcGF0aW5nIGRlbGl2ZXJpZXMgb2YgdGhlIGJvb2sgKipHZW5lc2lzIEVkaXRpb24gaW4gUTEgMjAyMS4qKlxuXG5BcyBhIGxpbWl0ZWQgcnVuIGZpcnN0IHByaW50LCB0aGUgYm9vayB3aWxsIGJlIGEgaGlnaCBxdWFsaXR5IGNvbGxlY3RvcnMgaXRlbS4gV2UgYXJlIHBheWluZyBjbG9zZSBhdHRlbnRpb24gdG8gZXZlcnkgYXNwZWN0LCBmcm9tIHRoZSBvdmVyYWxsIGRlc2lnbiwgdG8gdGhlIGZlYXR1cmVzIGhpZ2hsaWdodGluZyB0aGUgdW5pcXVlbmVzcyBvZiBlYWNoIGNvcHkuIEZvciBleGFtcGxlLCB0aGUgYnVybiB0cmFuc2FjdGlvbiBoYXNoIG9mIHJlZGVtcHRpb24gYmVpbmcgbWFya2VkIG9uIHRoZSBwaHlzaWNhbCBwcmludC4gXG5cbiMjIyBXaHkgYSBCb29rP1xuXG5XZSB0aGluayB0aGUgYXJ0aWZhY3QgdGhhdCBmcmFtZXMgdGhlIHdvcmsgY2FuIGJlIGFzIGltcG9ydGFudCBhcyB0aGUgY29udGVudCBpdHNlbGYuIFNvIG11Y2ggb2Ygd2hhdCB3ZSBjb25zdW1lIGV4aXN0cyBlcGhlbWVyYWxseSBpbiBmZWVkcyBhbmQgY2hhdHM6IHRoZSBCZWFjb24gQm9vayBjYXB0dXJlcyBhIHNuYXBzaG90IG9mIHNlbnRpbWVudCwgcG9ydGVkIGludG8gb2JqZWN0LWZvcm0uIFRoaXMgYXJ0aWZhY3Qgd2lsbCBsaXZlIG9uIGFzIHNvbWV0aGluZyB0byBiZSByZWFkLCByZWZlcmVuY2VkLCBhbmQgY2FyZWZ1bGx5IHN0b3JlZCBsb25nIGludG8gdGhlIGZ1dHVyZS5cblxuIyMjIEhvdyBjYW4gSSBnZXQgb25lP1xuXG5IZXJlIGFyZSB0aGUgZ2VuZXJhbCBkZXRhaWxzLCBzdWJqZWN0IHRvIGNoYW5nZS4gXG5cbioqTk9URTogVGhlcmUgaXMgbm8gd2F5IHRvIGFjcXVpcmUgdGhlIGJvb2sgYXQgdGhpcyB0aW1lLiBUaGUgZGF0ZSBhbmQgcmVsZWFzZSB3aWxsIGJlIHdlbGwgcHVibGljaXplZCBoZXJlIGFuZCB3aXRoIG91ciBwYXJ0bmVyIG9yZ3MuKipcblxuVGhlIGJ1bGsgb2YgdGhlIEdlbmVzaXMgRWRpdGlvbiAofjEwMCBib29rcykgd2lsbCBiZSBkaXN0cmlidXRlZCBvbi1jaGFpbiwgc2ltaWxhciB0byBVbmlzb2Nrcy4gVGhlIG1haW4gZGlmZmVyZW5jZSBpcyB0aGF0IHdlIHBsYW4gdG8gdXNlIGEgQmFsYW5jZXIgU21hcnQgcG9vbCB0byBhbGxvdyBtYXhpbWFsIGFjY2VzcyBmb3IgdGhlIGVudGlyZSBFdGhlcmV1bSBjb21tdW5pdHkuIExlYXJuIG1vcmUgYWJvdXQgaG93IHRoZXNlIG1lY2hhbmlzbXMgZnVuY3Rpb24gdGhyb3VnaCBjYXNlIHN0dWRpZXMgbGlrZSB0aGUgW0JhbmtsZXNzIEJBUF0oaHR0cHM6Ly9iYW5rbGVzcy5zdWJzdGFjay5jb20vcC90aGUtdWx0aW1hdGUtZ3VpZGUtdG8tYmFsYW5jZXItc21hcnQpIGFuZCBbUEVSUF0oaHR0cHM6Ly9wZXJwZXR1YWxwcm90b2NvbC5tZWRpdW0uY29tL3doeS13ZS1jaG9zZS10by1kaXN0cmlidXRlLXBlcnAtdXNpbmctYS1iYWxhbmNlci1saXF1aWRpdHktYm9vdHN0cmFwcGluZy1wb29sLWFhYzdmMWFiNjE4MSkuIFxuXG5BIGZldyBvZiB0aGUgYmVuZWZpdHMgaW5jbHVkZTpcblxuLSBiZXR0ZXIgcHJpY2UgZGlzY292ZXJ5OiBzdGFydHMgaGlnaCA-IGxvdywgbWFraW5nIGl0IGhhcmRlciBmb3Igd2hhbGVzIHRvIGNvcm5lciBzdXBwbHkgZWFybHkgb25cbi0gZmFpciBkaXN0cmlidXRpb24gZm9yIGh1bWFucyAoYm90cyBoYXZlIG5vIGFkdmFudGFnZSlcbi0gZWZmaWNpZW50IGZvciBib290c3RyYXBwaW5nIGxpcXVpZGl0eSBmb3IgdGhlIHBvb2wgY3JlYXRvclxuXG5BIHBvcnRpb24gb2YgdGhlIEdlbmVzaXMgRWRpdGlvbiAofjEwIGJvb2tzKSB3aWxsIGJlIHBsZWRnZWQgdG8gcGFydG5lciBjb21tdW5pdHkgb3JncyB0byBnaXZlIGF3YXkgdG8gdGhlaXIgY29tbXVuaXRpZXMuIEV0aGVyZXVtIGN1bHR1cmUgaXMgbm90aGluZyBpZiBub3QgcGFydGljaXBhdG9yeSwgYW5kIHdlIHdhbnQgdG8gZnVsbHkgZW1icmFjZSB0aGlzIGFzcGVjdC4gV2Ugd2lsbCBiZSBhbm5vdW5jaW5nIG91ciBwYXJ0bmVycyBjbG9zZXIgdG8gdGhlIGNvbXBsZXRlIGxhdW5jaC5cblxuVGhlcmUgd2lsbCBiZSBvdGhlciBmdW4gd2F5cyB0byB3aW4gYm9va3MsIHN0YXkgdHVuZWQgZm9yIG1vcmUgZGV0YWlscyEgSWYgeW91IGVuZCB1cCBub3QgZ2V0dGluZyBhIGJvb2ssIGRvbid0IHdvcnJ5OiBkaWdpdGFsIHZlcnNpb25zIG9mIHRoZSBjb250ZW50IHdpbGwgYmUgYXZhaWxhYmxlIGZvciBmcmVlIHRvIGVuam95LCBoZXJlIG9uIHRoaXMgYmxvZy5cblxuRmluYWxseSwgYW4gaW1wb3J0YW50IGdvYWwgb2YgdGhpcyBwcm9qZWN0IGlzIHRvIHByb21vdGUgYW5kIHN1cHBvcnQgRXRoZXJldW0gUHVibGljIEdvb2RzLiBUaGVyZWZvcmUsIGFueSBuZXQgRVRIIGdlbmVyYXRlZCBieSB0aGUgR2VuZXNpcyBzYWxlIG5ldCBvZiBwcmludGluZywgc2hpcHBpbmcgYW5kIGNvbnRyaWJ1dG9yIGNvc3RzIHdpbGwgYmUgZG9uYXRlZCB0byB0aGUgW09wZW4gR3JhbnRzIEV0aDIgc3RyZWFtXShodHRwczovL29wZW5ncmFudHMuY29tL2dyYW50LzB4NTNlN2RhYThlM2FhMjNjZDMwYzc1YjJmNTk5YzMwM2JhZGExNzA2NCkuXG5cbi0tLVxuXG5Zb3UgbWF5IGhhdmUgbm90aWNlZCB0aGlzIHBvc3QgaXMgaG9zdGVkIG9uIFtNaXJyb3IueHl6XShodHRwczovL21pcnJvci54eXovKSwgYSBuZXcgd2ViMyBwcm90b2NvbCBmb3IgcHVibGlzaGluZyB0byBjb21tdW5pdGllcy4gV2UncmUgZXhjaXRlZCB0byBiZSBvbmUgb2YgdGhlIGZpcnN0IHByb2plY3RzIHRvIHB1Ymxpc2ggdXNpbmcgaXQhIENoZWNrIGJhY2sgbmV4dCB3ZWVrIGZvciBtb3JlIGluZm8gb24gd2hhdCBTdGF0ZWZ1bCBXb3JrcyBpcyBhbGwgYWJvdXQgLSB3ZSBjYW4ndCB3YWl0IHRvIHNoYXJlIG1vcmUgYWJvdXQgb3VyIHZpc2lvbi5cblxuTWVzc2FnZSB1cyBvbiBUZWxlZ3JhbSBvciBUd2l0dGVyIFtAc3RhdGVmdWx3b3Jrc10oaHR0cHM6Ly90d2l0dGVyLmNvbS9TdGF0ZWZ1bFdvcmtzKSB3aXRoIHF1ZXN0aW9ucy4g4p2k77iPIiwidGltZXN0YW1wIjoxNjA4NzY2MDcwLCJ0aXRsZSI6IldlJ3JlIG1ha2luZyBhIGJvb2sgYWJvdXQgRXRoMiEiLCJwdWJsaWNhdGlvbiI6InN0YXRlZnVsIn0sImRpZ2VzdCI6ImExNTFlZTFkZWNiMjAyOGE4YmI0ODI3N2Y2OTI4YzZmMzgzMTljMzI2MDFkYzFkYTFlZTgyYWNmY2FkMmU1MjUiLCJhdXRob3JzaGlwIjp7ImNvbnRyaWJ1dG9yIjoiMHg0QzBhNDY2REYwNjI4RkU4Njk5MDUxYjNBYzY1MDY2NTMxOTFjYzIxIiwic2lnbmluZ0tleSI6Ii0tLS0tQkVHSU4gRUMgUFVCTElDIEtFWS0tLS0tXHJcbk1Ga3dFd1lIS29aSXpqMENBUVlJS29aSXpqMERBUWNEUWdBRWNMKytaNHdpdDFOWTYrb1VkaDc5bVFydHdzaVVcclxuQktVdXg5MHRpcGh6S0I0Zk9iWm1oK1pQYkJQNGhia3hCUGhHMzllMU5JeWdRZmY3Skc0YS9wNDczZz09XHJcbi0tLS0tRU5EIEVDIFBVQkxJQyBLRVktLS0tLVxyXG4iLCJzaWduYXR1cmUiOiI4akNxclJmZUxTNTZkbGhjaG10NUU2K0ltSTJhMWpBTzR6SzNtMThudkR3Q1N0K1ZOSlJ1dS8yQVJwQVpoRjlBd0F3OHo5MzlMaHp1d0dBVnZwRXVBdz09Iiwic2lnbmluZ0tleVNpZ25hdHVyZSI6IjB4ZjJjOWM0N2I1MjA2ZGIzNTQyOWMyNmYxOTVhMzJkMjk3OGM5YTU5Njc2ZjVlZmU5YWQyMzgxMDc5MDg3ODMwZTQ2ODU4YjMzY2NkMzNmN2M3MDhmZmFlMTc2NDk0YjBkMDhkZTNhOTVkMDU1YTg0NjE0MTE1ZDNhZmU0ZDlmYjgxYyIsInNpZ25pbmdLZXlNZXNzYWdlIjoiSSBhdXRob3JpemUgdGhlIHB1YmxpY2F0aW9uIG9mIGFydGljbGVzIG9uIHN0YXRlZnVsLm1pcnJvci54eXogZnJvbSB0aGlzIGRldmljZSB1c2luZzpcbi0tLS0tQkVHSU4gRUMgUFVCTElDIEtFWS0tLS0tXHJcbk1Ga3dFd1lIS29aSXpqMENBUVlJS29aSXpqMERBUWNEUWdBRWNMKytaNHdpdDFOWTYrb1VkaDc5bVFydHdzaVVcclxuQktVdXg5MHRpcGh6S0I0Zk9iWm1oK1pQYkJQNGhia3hCUGhHMzllMU5JeWdRZmY3Skc0YS9wNDczZz09XHJcbi0tLS0tRU5EIEVDIFBVQkxJQyBLRVktLS0tLVxyXG4iLCJhbGdvcml0aG0iOnsibmFtZSI6IkVDRFNBIiwiaGFzaCI6IlNIQS0yNTYifX0sInZlcnNpb24iOiIxMi0yMS0yMDIwIn0",
					},
				},
			},
			want: &schema.Feed{
				ID:       "lW0AMDN2RgOeqULk-u6Tv0wfZWpx9MfkrmqQQU-Mvuo",
				Chain:    filter.ChainArweaveMainnet,
				Index:    0,
				From:     "Ky1c1Kkt-jZ9sY1hvLF5nCf6WWdBhIU5Un_BMYh-t3c",
				To:       "Ky1c1Kkt-jZ9sY1hvLF5nCf6WWdBhIU5Un_BMYh-t3c",
				Type:     filter.TypeSocialPost,
				Platform: lo.ToPtr(filter.PlatformMirror),
				Fee: schema.Fee{
					Amount:  decimal.NewFromInt(27798031),
					Decimal: 12,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialPost,
						Tag:      filter.TagSocial,
						Platform: filter.PlatformMirror.String(),
						From:     "0x4C0a466DF0628FE8699051b3Ac6506653191cc21",
						To:       "Ky1c1Kkt-jZ9sY1hvLF5nCf6WWdBhIU5Un_BMYh-t3c",
						Metadata: &metadata.SocialPost{
							Title:         "We're making a book about Eth2!",
							Body:          "Stateful Works is excited to publicly introduce our first project, a limited run print book focused on the people and teams that brought the Ethereum Beacon Chain to life.\n\n### Who is Stateful Works?\n\nWe are a group of passionate longtime community members interested in exploring and creating Ethereum culture: reach out to [@statefulworks](https://twitter.com/StatefulWorks) on Telegram or Twitter with questions.\n\n### Book Details\n\nThis book will feature profiles, reflections, and stories from Eth2 implementers and researchers. We want to share the stories of the people that made this possible. We have been working on this for the past few months, and have made significant progress towards defining and realizing the vision.\n\nSo far, we have gathered profiles from over **25 Eth2 contributors**, with many more to come. We're anticipating deliveries of the book **Genesis Edition in Q1 2021.**\n\nAs a limited run first print, the book will be a high quality collectors item. We are paying close attention to every aspect, from the overall design, to the features highlighting the uniqueness of each copy. For example, the burn transaction hash of redemption being marked on the physical print. \n\n### Why a Book?\n\nWe think the artifact that frames the work can be as important as the content itself. So much of what we consume exists ephemerally in feeds and chats: the Beacon Book captures a snapshot of sentiment, ported into object-form. This artifact will live on as something to be read, referenced, and carefully stored long into the future.\n\n### How can I get one?\n\nHere are the general details, subject to change. \n\n**NOTE: There is no way to acquire the book at this time. The date and release will be well publicized here and with our partner orgs.**\n\nThe bulk of the Genesis Edition (~100 books) will be distributed on-chain, similar to Unisocks. The main difference is that we plan to use a Balancer Smart pool to allow maximal access for the entire Ethereum community. Learn more about how these mechanisms function through case studies like the [Bankless BAP](https://bankless.substack.com/p/the-ultimate-guide-to-balancer-smart) and [PERP](https://perpetualprotocol.medium.com/why-we-chose-to-distribute-perp-using-a-balancer-liquidity-bootstrapping-pool-aac7f1ab6181). \n\nA few of the benefits include:\n\n- better price discovery: starts high > low, making it harder for whales to corner supply early on\n- fair distribution for humans (bots have no advantage)\n- efficient for bootstrapping liquidity for the pool creator\n\nA portion of the Genesis Edition (~10 books) will be pledged to partner community orgs to give away to their communities. Ethereum culture is nothing if not participatory, and we want to fully embrace this aspect. We will be announcing our partners closer to the complete launch.\n\nThere will be other fun ways to win books, stay tuned for more details! If you end up not getting a book, don't worry: digital versions of the content will be available for free to enjoy, here on this blog.\n\nFinally, an important goal of this project is to promote and support Ethereum Public Goods. Therefore, any net ETH generated by the Genesis sale net of printing, shipping and contributor costs will be donated to the [Open Grants Eth2 stream](https://opengrants.com/grant/0x53e7daa8e3aa23cd30c75b2f599c303bada17064).\n\n---\n\nYou may have noticed this post is hosted on [Mirror.xyz](https://mirror.xyz/), a new web3 protocol for publishing to communities. We're excited to be one of the first projects to publish using it! Check back next week for more info on what Stateful Works is all about - we can't wait to share more about our vision.\n\nMessage us on Telegram or Twitter [@statefulworks](https://twitter.com/StatefulWorks) with questions. ❤️",
							PublicationID: "a151ee1decb2028a8bb48277f6928c6f38319c32601dc1da1ee82acfcad2e525",
							ContentURI:    "ar://lW0AMDN2RgOeqULk-u6Tv0wfZWpx9MfkrmqQQU-Mvuo",
							Target:        nil,
						},
					},
				},
				Status:    true,
				Timestamp: 1608766070,
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

			matched, err := instance.Match(ctx, testcase.arguments.task)
			testcase.wantError(t, err)
			require.True(t, matched)

			feed, err := instance.Transform(ctx, testcase.arguments.task)
			testcase.wantError(t, err)

			require.Equal(t, testcase.want, feed)

			t.Log(feed)
		})
	}
}
