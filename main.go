package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// Envelope was generated 2024-08-29 10:38:29 by https://xml-to-go.github.io/ in Ukraine.
type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	S       string   `xml:"s,attr"`
	A       string   `xml:"a,attr"`
	U       string   `xml:"u,attr"`
	Header  struct {
		Text   string `xml:",chardata"`
		Action struct {
			Text           string `xml:",chardata"`
			MustUnderstand string `xml:"mustUnderstand,attr"`
		} `xml:"Action"`
		RelatesTo string `xml:"RelatesTo"`
		Security  struct {
			Text           string `xml:",chardata"`
			MustUnderstand string `xml:"mustUnderstand,attr"`
			O              string `xml:"o,attr"`
			Timestamp      struct {
				Text    string `xml:",chardata"`
				ID      string `xml:"Id,attr"`
				Created string `xml:"Created"`
				Expires string `xml:"Expires"`
			} `xml:"Timestamp"`
		} `xml:"Security"`
	} `xml:"Header"`
	Body struct {
		Text                string `xml:",chardata"`
		Xsi                 string `xml:"xsi,attr"`
		Xsd                 string `xml:"xsd,attr"`
		SyncUpdatesResponse struct {
			Text              string `xml:",chardata"`
			Xmlns             string `xml:"xmlns,attr"`
			SyncUpdatesResult struct {
				Text       string `xml:",chardata"`
				NewUpdates struct {
					Text       string `xml:",chardata"`
					UpdateInfo []struct {
						Text       string `xml:",chardata"`
						ID         string `xml:"ID"`
						Deployment struct {
							Text                  string `xml:",chardata"`
							ID                    string `xml:"ID"`
							Action                string `xml:"Action"`
							IsAssigned            string `xml:"IsAssigned"`
							LastChangeTime        string `xml:"LastChangeTime"`
							AutoSelect            string `xml:"AutoSelect"`
							AutoDownload          string `xml:"AutoDownload"`
							SupersedenceBehavior  string `xml:"SupersedenceBehavior"`
							Priority              string `xml:"Priority"`
							HandlerSpecificAction string `xml:"HandlerSpecificAction"`
							FlightId              string `xml:"FlightId"`
						} `xml:"Deployment"`
						IsLeaf string `xml:"IsLeaf"`
						XML    struct {
							Text           string `xml:",chardata"`
							UpdateIdentity struct {
								Text           string `xml:",chardata"`
								UpdateID       string `xml:"UpdateID,attr"`
								RevisionNumber string `xml:"RevisionNumber,attr"`
							} `xml:"UpdateIdentity"`
							Properties struct {
								Text                 string `xml:",chardata"`
								UpdateType           string `xml:"UpdateType,attr"`
								PerUser              string `xml:"PerUser,attr"`
								PackageRank          string `xml:"PackageRank,attr"`
								IsAppxFramework      string `xml:"IsAppxFramework,attr"`
								ExplicitlyDeployable string `xml:"ExplicitlyDeployable,attr"`
								ApplyPackageRank     string `xml:"ApplyPackageRank,attr"`
								SecuredFragment      string `xml:"SecuredFragment"`
							} `xml:"Properties"`
							Relationships struct {
								Text          string `xml:",chardata"`
								Prerequisites struct {
									Text       string `xml:",chardata"`
									AtLeastOne []struct {
										Text           string `xml:",chardata"`
										IsCategory     string `xml:"IsCategory,attr"`
										UpdateIdentity []struct {
											Text     string `xml:",chardata"`
											UpdateID string `xml:"UpdateID,attr"`
										} `xml:"UpdateIdentity"`
									} `xml:"AtLeastOne"`
									UpdateIdentity []struct {
										Text     string `xml:",chardata"`
										UpdateID string `xml:"UpdateID,attr"`
									} `xml:"UpdateIdentity"`
								} `xml:"Prerequisites"`
								BundledUpdates struct {
									Text       string `xml:",chardata"`
									AtLeastOne struct {
										Text           string `xml:",chardata"`
										UpdateIdentity []struct {
											Text           string `xml:",chardata"`
											UpdateID       string `xml:"UpdateID,attr"`
											RevisionNumber string `xml:"RevisionNumber,attr"`
										} `xml:"UpdateIdentity"`
									} `xml:"AtLeastOne"`
									UpdateIdentity []struct {
										Text           string `xml:",chardata"`
										UpdateID       string `xml:"UpdateID,attr"`
										RevisionNumber string `xml:"RevisionNumber,attr"`
									} `xml:"UpdateIdentity"`
								} `xml:"BundledUpdates"`
							} `xml:"Relationships"`
							ApplicabilityRules struct {
								Text        string `xml:",chardata"`
								IsInstalled struct {
									Text                 string `xml:",chardata"`
									AppxPackageInstalled string `xml:"AppxPackageInstalled"`
									True                 string `xml:"True"`
								} `xml:"IsInstalled"`
								Metadata struct {
									Text                string `xml:",chardata"`
									AppxPackageMetadata struct {
										Text               string `xml:",chardata"`
										AppxFamilyMetadata struct {
											Text                  string `xml:",chardata"`
											Name                  string `xml:"Name,attr"`
											Publisher             string `xml:"Publisher,attr"`
											LegacyMobileProductId string `xml:"LegacyMobileProductId,attr"`
										} `xml:"AppxFamilyMetadata"`
										AppxMetadata struct {
											Text              string `xml:",chardata"`
											PackageType       string `xml:"PackageType,attr"`
											IsAppxBundle      string `xml:"IsAppxBundle,attr"`
											PackageMoniker    string `xml:"PackageMoniker,attr"`
											ApplicabilityBlob string `xml:"ApplicabilityBlob"`
										} `xml:"AppxMetadata"`
									} `xml:"AppxPackageMetadata"`
								} `xml:"Metadata"`
								IsInstallable struct {
									Text                   string `xml:",chardata"`
									AppxPackageInstallable string `xml:"AppxPackageInstallable"`
								} `xml:"IsInstallable"`
							} `xml:"ApplicabilityRules"`
						} `xml:"Xml"`
						IsShared string `xml:"IsShared"`
					} `xml:"UpdateInfo"`
				} `xml:"NewUpdates"`
				Truncated string `xml:"Truncated"`
				NewCookie struct {
					Text          string `xml:",chardata"`
					Expiration    string `xml:"Expiration"`
					EncryptedData string `xml:"EncryptedData"`
				} `xml:"NewCookie"`
				DeployedOutOfScopeRevisionIds struct {
					Text string   `xml:",chardata"`
					Int  []string `xml:"int"`
				} `xml:"DeployedOutOfScopeRevisionIds"`
				DriverSyncNotNeeded string `xml:"DriverSyncNotNeeded"`
				ExtendedUpdateInfo  struct {
					Text    string `xml:",chardata"`
					Updates struct {
						Text   string `xml:",chardata"`
						Update []struct {
							Text string `xml:",chardata"`
							ID   string `xml:"ID"`
							XML  struct {
								Text               string `xml:",chardata"`
								ExtendedProperties struct {
									Text                      string `xml:",chardata"`
									DefaultPropertiesLanguage string `xml:"DefaultPropertiesLanguage,attr"`
									CreationDate              string `xml:"CreationDate,attr"`
									ContentType               string `xml:"ContentType,attr"`
									IsAppxFramework           string `xml:"IsAppxFramework,attr"`
									FromStoreService          string `xml:"FromStoreService,attr"`
									PackageIdentityName       string `xml:"PackageIdentityName,attr"`
									LegacyMobileProductId     string `xml:"LegacyMobileProductId,attr"`
									Handler                   string `xml:"Handler,attr"`
									MaxDownloadSize           string `xml:"MaxDownloadSize,attr"`
									MinDownloadSize           string `xml:"MinDownloadSize,attr"`
									MandatoryVersion          string `xml:"MandatoryVersion,attr"`
									MandatoryDate             string `xml:"MandatoryDate,attr"`
									PackageContentId          string `xml:"PackageContentId,attr"`
									InstallationBehavior      string `xml:"InstallationBehavior"`
								} `xml:"ExtendedProperties"`
								Files struct {
									Text string `xml:",chardata"`
									File []struct {
										Text                        string `xml:",chardata"`
										FileName                    string `xml:"FileName,attr"`
										Digest                      string `xml:"Digest,attr"`
										DigestAlgorithm             string `xml:"DigestAlgorithm,attr"`
										Size                        string `xml:"Size,attr"`
										Modified                    string `xml:"Modified,attr"`
										InstallerSpecificIdentifier string `xml:"InstallerSpecificIdentifier,attr"`
										PatchingType                string `xml:"PatchingType,attr"`
										AdditionalDigest            struct {
											Text      string `xml:",chardata"`
											Algorithm string `xml:"Algorithm,attr"`
										} `xml:"AdditionalDigest"`
										PiecesHashDigest struct {
											Text      string `xml:",chardata"`
											Algorithm string `xml:"Algorithm,attr"`
										} `xml:"PiecesHashDigest"`
										BlockMapDigest struct {
											Text      string `xml:",chardata"`
											Algorithm string `xml:"Algorithm,attr"`
										} `xml:"BlockMapDigest"`
									} `xml:"File"`
								} `xml:"Files"`
								HandlerSpecificData struct {
									Text                   string `xml:",chardata"`
									Type                   string `xml:"type,attr"`
									AppxPackageInstallData struct {
										Text            string `xml:",chardata"`
										PackageFileName string `xml:"PackageFileName,attr"`
										MainPackage     string `xml:"MainPackage,attr"`
									} `xml:"AppxPackageInstallData"`
									CategoryInformation struct {
										Text                   string `xml:",chardata"`
										CategoryType           string `xml:"CategoryType,attr"`
										ProhibitsSubcategories string `xml:"ProhibitsSubcategories,attr"`
										ProhibitsUpdates       string `xml:"ProhibitsUpdates,attr"`
										ExcludedByDefault      string `xml:"ExcludedByDefault,attr"`
									} `xml:"CategoryInformation"`
								} `xml:"HandlerSpecificData"`
							} `xml:"Xml"`
						} `xml:"Update"`
					} `xml:"Updates"`
				} `xml:"ExtendedUpdateInfo"`
			} `xml:"SyncUpdatesResult"`
		} `xml:"SyncUpdatesResponse"`
	} `xml:"Body"`
}

// Envelope was generated 2024-08-29 10:17:07 by https://xml-to-go.github.io/ in Ukraine.
type Envelope2 struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	S       string   `xml:"s,attr"`
	A       string   `xml:"a,attr"`
	U       string   `xml:"u,attr"`
	Header  struct {
		Text   string `xml:",chardata"`
		Action struct {
			Text           string `xml:",chardata"`
			MustUnderstand string `xml:"mustUnderstand,attr"`
		} `xml:"Action"`
		RelatesTo string `xml:"RelatesTo"`
		Security  struct {
			Text           string `xml:",chardata"`
			MustUnderstand string `xml:"mustUnderstand,attr"`
			O              string `xml:"o,attr"`
			Timestamp      struct {
				Text    string `xml:",chardata"`
				ID      string `xml:"Id,attr"`
				Created string `xml:"Created"`
				Expires string `xml:"Expires"`
			} `xml:"Timestamp"`
		} `xml:"Security"`
	} `xml:"Header"`
	Body struct {
		Text              string `xml:",chardata"`
		Xsi               string `xml:"xsi,attr"`
		Xsd               string `xml:"xsd,attr"`
		GetCookieResponse struct {
			Text            string `xml:",chardata"`
			Xmlns           string `xml:"xmlns,attr"`
			GetCookieResult struct {
				Text          string `xml:",chardata"`
				Expiration    string `xml:"Expiration"`
				EncryptedData string `xml:"EncryptedData"`
			} `xml:"GetCookieResult"`
		} `xml:"GetCookieResponse"`
	} `xml:"Body"`
}

var (
	catId          = "d25480ca-36aa-46e6-b76b-39608d49558c" // mc category id
	Retail         = "Retail"
	ReleasePreview = "RP"
	InsiderSlow    = "WIS"
	InsiderFast    = "WIF"
)

func GetCookie() string {
	xmlRequestBody, err := os.ReadFile("GetCookie.xml")
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.Post("https://fe3.delivery.mp.microsoft.com/ClientWebService/client.asmx", "application/soap+xml", bytes.NewReader(xmlRequestBody))
	if err != nil {
		log.Fatal(err)
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data Envelope2
	err = xml.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data.Body.GetCookieResponse.GetCookieResult.EncryptedData
}

func main() {
	cookie := GetCookie()

	xmlTemplate, err := os.ReadFile("WUIDRequest.xml")
	if err != nil {
		log.Fatal(err)
	}

	xmlContent := string(xmlTemplate)
	xmlContent = strings.Replace(xmlContent, "{0}", cookie, 1)
	xmlContent = strings.Replace(xmlContent, "{1}", catId, 1)
	xmlContent = strings.Replace(xmlContent, "{2}", Retail, 1)

	req, err := http.Post("https://fe3.delivery.mp.microsoft.com/ClientWebService/client.asmx", "application/soap+xml", strings.NewReader(xmlContent))
	if err != nil {
		log.Fatal(err)
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}

	content := html.UnescapeString(string(body))

	var data Envelope
	err = xml.Unmarshal([]byte(content), &data)
	if err != nil {
		log.Fatal(err)
	}

	for _, a := range data.Body.SyncUpdatesResponse.SyncUpdatesResult.NewUpdates.UpdateInfo {
		if strings.Contains(a.XML.ApplicabilityRules.Metadata.AppxPackageMetadata.AppxMetadata.PackageMoniker, "Microsoft.Services.Store.Engagement_") && strings.Contains(a.XML.ApplicabilityRules.Metadata.AppxPackageMetadata.AppxMetadata.PackageMoniker, "x64") {
			fmt.Println("Update id for latest x64 MC bedrock: " + a.XML.UpdateIdentity.UpdateID)
		}
	}

	//fmt.Println(content)
}
