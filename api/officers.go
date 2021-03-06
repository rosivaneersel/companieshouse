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
)

type (
	// DateOfBirth struct consists of Day(int), Month (int) and Year (int)
	OfficerDateOfBirth struct {
		Day   int `json:"day"`
		Month int `json:"month"`
		Year  int `json:"year"`
	}

	// Identification struct
	Identification struct {
		IDType             string `json:"identification_type"`
		Authority          string `json:"legal_authority"`
		LegalForm          string `json:"legal_form"`
		PlaceRegistered    string `json:"place_registered"`
		RegistrationNumber string `json:"registration_number"`
	}

	// Officer struct contains the data of a company's officers
	Officer struct {
		Address            Address            `json:"address"`
		AppointedOn        ChDate             `json:"appointed_on"`
		CountryOfResidence string             `json:"country_of_residence"`
		DateOfBirth        OfficerDateOfBirth `json:"date_of_birth"`
		FormerNames        []struct {
			Forenames string `json:"forenames"`
			Surname   string `json:"surname"`
		} `json:"former_names"`
		Identification Identification `json:"identification"`
		Links          struct {
			Officer struct {
				Appointments string `json:"appointments"`
			} `json:"officer"`
		} `json:"links"`
		Name        string `json:"name"`
		Nationality string `json:"nationality"`
		Occupation  string `json:"occupation"`
		OfficerRole string `json:"officer_role"`
		ResignedOn  ChDate `json:"resigned_on"`
	}

	// OfficerResponse contains the server response of a data request to the companies house API
	OfficerResponse struct {
		Etag          string    `json:"etag"`
		Kind          string    `json:"kind"`
		Start         int       `json:"start_index"`
		ItemsPerPage  int       `json:"items_per_page"`
		TotalResults  int       `json:"total_results"`
		ActiveCount   int       `json:"active_count"`
		InactiveCount int       `json:"inactive_count"`
		ResignedCount int       `json:"resigned_count"`
		Items         []*Officer `json:"items"`
		Links         struct {
			self string `json:"self"`
		} `json:"Links"`
	}
)

// GetOfficers gets the json data for a company's officers from the Companies House REST API
// and returns a new OfficersResponse and an error
func (c *Company) getOfficers() (*OfficerResponse, error) {
	officers := &OfficerResponse{}
	resp, err := c.api.CallAPI("/company/"+c.CompanyNumber+"/officers", nil, false, ContentTypeJSON)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &officers)
	if err != nil {
		return nil, err
	}

	return officers, nil
}

func (a *API) GetOfficers(c string) (<-chan *OfficerResponse, <-chan error) {
	r := make(chan *OfficerResponse, 1)
	e := make(chan error, 1)

	go func() {
		defer close(r)
		defer close(e)

		officers := &OfficerResponse{}
			resp, err := a.CallAPI("/company/" + c +"/officers", nil, false, ContentTypeJSON)
		if err != nil {
			r <- nil
			e <- err
			return
		}

		err = json.Unmarshal(resp, &officers)
		if err != nil {
			r <- nil
			e <- err
			return
		}
		r <- officers
		e <- nil
	}()

	return r, e
}
