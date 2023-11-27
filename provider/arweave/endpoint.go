package arweave

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=Endpoint --linecomment --output endpoint_string.go
type Endpoint int

const (
	EndpointArweave         Endpoint = iota + 1 // https://arweave.net/
	EndpointArweaveFllstck                      // https://arweave.fllstck.dev/
	EndpointARIO                                // https://ar-io.dev/
	EndpointPermagate                           // https://permagate.io/
	EndpointLove4Src                            // https://love4src.com/
	EndpointBobInstein                          // https://bobinstein.com/
	EndpointPie                                 // https://gatewaypie.com/
	EndpointAleko0o                             //https://aleko0o.store/
	EndpointVevivo                              //https://vevivo.xyz/
	EndpointSulapan                             //https://sulapan.com/
	EndpointBicem                               //https://bicem.xyz/
	EndpointDilsinay                            //https://dilsinay.online/
	EndpointLostgame                            //https://lostgame.online/
	EndpointKhaldrogo                           //https://khaldrogo.site/
	EndpointDasamuka                            //https://dasamuka.cloud/
	EndpointArendor                             //https://arendor.xyz/
	EndpointVelaryon                            //https://velaryon.xyz/
	EndpointKingsharald                         //https://kingsharald.xyz/
	EndpointMoruehoca                           //https://moruehoca.online/
	EndpointKazazel                             //https://kazazel.xyz/
	EndpointAhmkah                              //https://ahmkah.online/
	EndpointElessar                             //https://elessardarken.xyz/
	EndpointYusufaytnxyzWay                     //https://yusufaytn.xyz/
	EndpointarAlper                             //https://aralper.xyz/
	EndpointGren                                //https://grenimo.xyz/
	EndpointVikanren                            //https://vikanren.xyz/
	EndpointGateOfRivia                         //https://neuweltgeld.xyz/
	EndpointMegumii                             //https://blessingway.xyz/
	EndpointRedwhite                            //https://redwhiteconnect.xyz/
	EndpointBayurnet                            //https://fisneci.com/
)
