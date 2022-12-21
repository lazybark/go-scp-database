package wikidot

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lazybark/go-scp-database/pkg/scp"
	"github.com/pterm/pterm"
)

var ratingRegex = regexp.MustCompile(`<span class="rate-points">rating:&nbsp;<span class="number prw54353">(.*?)</span></span>`)
var hasNewClasses = regexp.MustCompile(`class="main-class"`)
var objectClassOld = regexp.MustCompile(`<p><strong>Object Class:</strong> (.*?)</p>`)
var objectClassNew = regexp.MustCompile("<div class=\"contain-class\">\n<div class=\"class-category\">Containment Class:</div>\n<div class=\"class-text\">(.*?)</div>\n</div>")
var objectClassSec = regexp.MustCompile("<div class=\"second-class\">\n<div class=\"class-category\">Secondary Class:</div>\n<div class=\"class-text\">(.*?)</div>\n</div>")
var objectClassDisrupt = regexp.MustCompile("<div class=\"disrupt-class\">\n<div class=\"class-category\">Disruption Class:</div>\n<div class=\"class-text\">(.*?)</div>\n</div>")
var objectClassRisk = regexp.MustCompile("<div class=\"risk-class\">\n<div class=\"class-category\">Risk Class:</div>\n<div class=\"class-text\">(.*?)</div>\n</div>")

var textOld = regexp.MustCompile("<p><strong>Object Class:</strong> (.*)</p>(.*)<div class=\"footnotes-footer\">\n<div class=\"title\">Footnotes</div>")

type WikiDot struct {
	site string
}

func NewDB() scp.ISCPDatabase {
	return WikiDot{site: "https://scp-wiki.wikidot.com"}
}

func (w WikiDot) Connect() error {
	spinnerInfo, _ := pterm.DefaultSpinner.Start("Connecting to SCP network...")
	time.Sleep(time.Second * 1)
	spinnerInfo.UpdateText(fmt.Sprintf("Connected to SCP network. Session ID %v", uuid.New().String()))
	time.Sleep(time.Second * 1)
	//spinnerInfo.UpdateText("Found free socket at server 172.17.2.53")
	spinnerInfo.Success()

	return nil
}
func (w WikiDot) GetObject(id string) (scp.ISCPObject, error) {
	spinnerInfo2, _ := pterm.DefaultSpinner.Start(fmt.Sprintf("Looking for object: SCP-%s", id))
	spinnerInfo2.RemoveWhenDone = true
	defer spinnerInfo2.Stop()

	o := WikiDotObj{
		rating: "Unknown",
		code:   "SCP-" + id,
	}

	url := "https://scp-wiki.wikidot.com" + "/scp-" + id
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	if resp.Status != "200 OK" {
		return nil, fmt.Errorf("object does not exist in Foundation database")
	}

	html := string(body)

	spinnerInfo2.UpdateText("Object found in Foundation database. Scanning...")

	//Set rating
	resultsRatingCheck := ratingRegex.FindStringSubmatch(html)
	if resultsRatingCheck != nil && resultsRatingCheck[1] != "" {
		o.rating = resultsRatingCheck[1]
	}

	//Set number
	newClassesCheck := hasNewClasses.FindStringSubmatch(html)
	if len(newClassesCheck) > 0 {
		o.hasNewClasses = true
	}

	//If we have old class system - just one check
	if !o.hasNewClasses {
		/*resultsOldClassCheck := objectClassOld.FindStringSubmatch(html)
		if resultsOldClassCheck != nil && resultsOldClassCheck[1] != "" {
			for cl, cn := range scp.ContainmentClasses {
				if strings.EqualFold(cn, resultsOldClassCheck[1]) {
					o.containmentClass = cl
				}
			}
		}*/

		resultsTextOld := textOld.FindAllString(html, 2)
		if resultsTextOld != nil && resultsTextOld[1] != "" {
			for cl, cn := range scp.ContainmentClasses {
				if strings.EqualFold(cn, resultsTextOld[1]) {
					o.containmentClass = cl

				}
			}
			if resultsTextOld[2] != "" {
				o.textOld = resultsTextOld[2]
			}
		}

	} else {
		//For new classes we have lot more checks
		resultsNewClassCheck := objectClassNew.FindStringSubmatch(html)
		if resultsNewClassCheck != nil && resultsNewClassCheck[1] != "" {
			for cl, cn := range scp.ContainmentClasses {
				if strings.EqualFold(cn, resultsNewClassCheck[1]) {
					o.containmentClass = cl
				}
			}
		}

		resultsClassSecCheck := objectClassSec.FindStringSubmatch(html)
		if resultsClassSecCheck != nil && resultsClassSecCheck[1] != "" {
			for cl, cn := range scp.EsotecicClasses {
				if strings.EqualFold(cn, resultsClassSecCheck[1]) {
					o.esotecicClass = scp.EsotecicClass(cl)
				}
			}
		}

		resultsClassDisruptCheck := objectClassDisrupt.FindStringSubmatch(html)
		if resultsClassDisruptCheck != nil && resultsClassDisruptCheck[1] != "" {
			for cl, cn := range scp.DisruptionClasses {
				if strings.EqualFold(cn, resultsClassDisruptCheck[1]) {
					o.disruptionClass = scp.DisruptionClass(cl)
				}
			}
		}

		resultsClassRiskCheck := objectClassRisk.FindStringSubmatch(html)
		if resultsClassRiskCheck != nil && resultsClassRiskCheck[1] != "" {
			for cl, cn := range scp.RiskClasses {
				if strings.EqualFold(cn, resultsClassRiskCheck[1]) {
					o.riskClass = scp.RiskClass(cl)
				}
			}
		}
	}
	/*

		number := ""
		if resultsNum != nil && resultsNum[1] != "" {
			number = resultsNum[1]
		} else {
			resultsNum = numberOldRegex.FindStringSubmatch(html)
			if resultsNum != nil && resultsNum[1] != "" {
				number = resultsNum[1]
			}
		}
		if number != "" {
			o.number, err = strconv.Atoi(number)
			if err != nil {
				return nil, err
			}
		}*/

	return o, nil
}
