/*
Golang Companies House REST service API
Copyright (C) 2017, Balkan C & T OOD

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package companieshouse

import (
	"encoding/json"
	"github.com/BalkanTech/companieshouse/api/testhandlers"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// TestGetCharges is a unit test of the GetCharges function
// The test loops over all companies in TestCompanies and then checks:
// 1) If a call can be made 2) If the response contains the expected content
func TestGetCharges(t *testing.T) {
	companyServer := httptest.NewServer(http.HandlerFunc(testhandlers.GetCompanyHandler))
	defer companyServer.Close()

	insolvenciesServer := httptest.NewServer(http.HandlerFunc(testhandlers.GetChargesHandler))
	defer insolvenciesServer.Close()
	api := API{}
	api.SetAPIKey("12345")

	t.Log("Testing the GetCharges function")
	{
		t.Log("\tWhen checking companies for charges")
		for _, test := range TestCompanies {
			api.setAPIURL(companyServer.URL)
			c, err := api.GetCompany(test.company)
			{
				if err != nil {
					t.Fatalf("\t\tShould be able to get a company \"%s\": %s. %v", test.company, err, testFailed)
				}
				t.Logf("\t\tShould be able to get a company \"%s\". %v", test.company, testOK)

				api.setAPIURL(insolvenciesServer.URL)

				data, _ := testhandlers.GetResponse(test.company, testhandlers.Charges)
				expected := &ChargesResponse{}
				json.Unmarshal([]byte(data), expected)

				t.Logf("\t\t\tWhile using company %s(%s)", c.CompanyName, c.CompanyNumber)
				resp, err := c.GetCharges()
				{
					if err != nil {
						t.Fatal("\t\t\t\tShould be able to get the charges.", testFailed, err)
					}
					t.Log("\t\t\t\tShould be able to get the charges. ", testOK)

					if reflect.DeepEqual(*resp, *expected) != true {
						t.Fatalf("\t\t\t\tShould have proper content, received %+v, but expected %+v. %v", resp, c, testFailed)
					}
					t.Log("\t\t\t\tShould have proper content. ", testOK)
				}
			}
		}
	}
}
