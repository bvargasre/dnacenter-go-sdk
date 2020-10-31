package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// SoftwareImageManagementSWIMService is the service to communicate with the SoftwareImageManagementSWIM API endpoint
type SoftwareImageManagementSWIMService service

// ActivateDTO is the ActivateDTO definition
type ActivateDTO struct {
	ActivateLowerImageVersion bool     `json:"activateLowerImageVersion,omitempty"` //
	DeviceUpgradeMode         string   `json:"deviceUpgradeMode,omitempty"`         //
	DeviceUuid                string   `json:"deviceUuid,omitempty"`                //
	DistributeIfNeeded        bool     `json:"distributeIfNeeded,omitempty"`        //
	ImageUuidList             []string `json:"imageUuidList,omitempty"`             //
	SmuImageUuidList          []string `json:"smuImageUuidList,omitempty"`          //
}

// DistributeDTO is the DistributeDTO definition
type DistributeDTO struct {
	DeviceUuid string `json:"deviceUuid,omitempty"` //
	ImageUuid  string `json:"imageUuid,omitempty"`  //
}

// ImageImportFromUrlDTO is the ImageImportFromUrlDTO definition
type ImageImportFromUrlDTO struct {
	ApplicationType string `json:"applicationType,omitempty"` //
	ImageFamily     string `json:"imageFamily,omitempty"`     //
	SourceURL       string `json:"sourceURL,omitempty"`       //
	ThirdParty      bool   `json:"thirdParty,omitempty"`      //
	Vendor          string `json:"vendor,omitempty"`          //
}

// ApplicableDevicesForImage is the ApplicableDevicesForImage definition
type ApplicableDevicesForImage struct {
	MdfId       string   `json:"mdfId,omitempty"`       //
	ProductId   []string `json:"productId,omitempty"`   //
	ProductName string   `json:"productName,omitempty"` //
}

// ImageInfoListResponse is the ImageInfoListResponse definition
type ImageInfoListResponse struct {
	Response []Response `json:"response,omitempty"` //
	Version  string     `json:"version,omitempty"`  //
}

// ProfileInfo is the ProfileInfo definition
type ProfileInfo struct {
	Description        string `json:"description,omitempty"`        //
	ExtendedAttributes string `json:"extendedAttributes,omitempty"` //
	Memory             int    `json:"memory,omitempty"`             //
	ProductType        string `json:"productType,omitempty"`        //
	ProfileName        string `json:"profileName,omitempty"`        //
	Shares             int    `json:"shares,omitempty"`             //
	VCpu               int    `json:"vCpu,omitempty"`               //
}

// Response is the Response definition
type Response struct {
	TaskId string `json:"taskId,omitempty"` //
	Url    string `json:"url,omitempty"`    //
}

// TaskIdResult is the TaskIdResult definition
type TaskIdResult struct {
	Response Response `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// GetSoftwareImageDetailsQueryParams defines the query parameters for this request
type GetSoftwareImageDetailsQueryParams struct {
	ImageUuid            string  `url:"imageUuid,omitempty"`            // imageUuid
	Name                 string  `url:"name,omitempty"`                 // name
	Family               string  `url:"family,omitempty"`               // family
	ApplicationType      string  `url:"applicationType,omitempty"`      // applicationType
	ImageIntegrityStatus string  `url:"imageIntegrityStatus,omitempty"` // imageIntegrityStatus - FAILURE, UNKNOWN, VERIFIED
	Version              string  `url:"version,omitempty"`              // software Image Version
	ImageSeries          string  `url:"imageSeries,omitempty"`          // image Series
	ImageName            string  `url:"imageName,omitempty"`            // image Name
	IsTaggedGolden       boolean `url:"isTaggedGolden,omitempty"`       // is Tagged Golden
	IsCCORecommended     boolean `url:"isCCORecommended,omitempty"`     // is recommended from cisco.com
	IsCCOLatest          boolean `url:"isCCOLatest,omitempty"`          // is latest from cisco.com
	CreatedTime          int     `url:"createdTime,omitempty"`          // time in milliseconds (epoch format)
	ImageSizeGreaterThan int     `url:"imageSizeGreaterThan,omitempty"` // size in bytes
	ImageSizeLesserThan  int     `url:"imageSizeLesserThan,omitempty"`  // size in bytes
	SortBy               string  `url:"sortBy,omitempty"`               // sort results by this field
	SortOrder            string  `url:"sortOrder,omitempty"`            // sort order - 'asc' or 'des'. Default is asc
	Limit                int     `url:"limit,omitempty"`                // limit
	Offset               int     `url:"offset,omitempty"`               // offset
}

// GetSoftwareImageDetails getSoftwareImageDetails
/* Returns software image list based on a filter criteria. For example: "filterbyName = cat3k%"
@param imageUuid imageUuid
@param name name
@param family family
@param applicationType applicationType
@param imageIntegrityStatus imageIntegrityStatus - FAILURE, UNKNOWN, VERIFIED
@param version software Image Version
@param imageSeries image Series
@param imageName image Name
@param isTaggedGolden is Tagged Golden
@param isCCORecommended is recommended from cisco.com
@param isCCOLatest is latest from cisco.com
@param createdTime time in milliseconds (epoch format)
@param imageSizeGreaterThan size in bytes
@param imageSizeLesserThan size in bytes
@param sortBy sort results by this field
@param sortOrder sort order - 'asc' or 'des'. Default is asc
@param limit limit
@param offset offset
*/
func (s *SoftwareImageManagementSWIMService) GetSoftwareImageDetails(getSoftwareImageDetailsQueryParams *GetSoftwareImageDetailsQueryParams) (*ImageInfoListResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/image/importation"

	queryString, _ := query.Values(getSoftwareImageDetailsQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&ImageInfoListResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*ImageInfoListResponse)
	return result, response, err

}

// ImportLocalSoftwareImageQueryParams defines the query parameters for this request
type ImportLocalSoftwareImageQueryParams struct {
	IsThirdParty              boolean `url:"isThirdParty,omitempty"`              // Third party Image check
	ThirdPartyVendor          string  `url:"thirdPartyVendor,omitempty"`          // Third Party Vendor
	ThirdPartyImageFamily     string  `url:"thirdPartyImageFamily,omitempty"`     // Third Party image family
	ThirdPartyApplicationType string  `url:"thirdPartyApplicationType,omitempty"` // Third Party Application Type
}

// ImportLocalSoftwareImage importLocalSoftwareImage
/* Fetches a software image from local file system and uploads to DNA Center. Supported software image files extensions are bin, img, tar, smu, pie, aes, iso, ova, tar_gz and qcow2
@param Content-Type Request body content type
@param isThirdParty Third party Image check
@param thirdPartyVendor Third Party Vendor
@param thirdPartyImageFamily Third Party image family
@param thirdPartyApplicationType Third Party Application Type
*/
func (s *SoftwareImageManagementSWIMService) ImportLocalSoftwareImage(importLocalSoftwareImageQueryParams *ImportLocalSoftwareImageQueryParams) (*TaskIdResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/image/importation/source/file"

	queryString, _ := query.Values(importLocalSoftwareImageQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&TaskIdResult{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TaskIdResult)
	return result, response, err

}

// ImportSoftwareImageViaURLQueryParams defines the query parameters for this request
type ImportSoftwareImageViaURLQueryParams struct {
	ScheduleAt     string `url:"scheduleAt,omitempty"`     // Epoch Time (The number of milli-seconds since January 1 1970 UTC) at which the distribution should be scheduled (Optional)
	ScheduleDesc   string `url:"scheduleDesc,omitempty"`   // Custom Description (Optional)
	ScheduleOrigin string `url:"scheduleOrigin,omitempty"` // Originator of this call (Optional)
}

// ImportSoftwareImageViaURL importSoftwareImageViaURL
/* Fetches a software image from remote file system (using URL for HTTP/FTP) and uploads to DNA Center. Supported image files extensions are bin, img, tar, smu, pie, aes, iso, ova, tar_gz and qcow2
@param scheduleAt Epoch Time (The number of milli-seconds since January 1 1970 UTC) at which the distribution should be scheduled (Optional)
@param scheduleDesc Custom Description (Optional)
@param scheduleOrigin Originator of this call (Optional)
*/
func (s *SoftwareImageManagementSWIMService) ImportSoftwareImageViaURL(importSoftwareImageViaURLQueryParams *ImportSoftwareImageViaURLQueryParams, importSoftwareImageViaURLRequest *ImportSoftwareImageViaURLRequest) (*TaskIdResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/image/importation/source/url"

	queryString, _ := query.Values(importSoftwareImageViaURLQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetBody(importSoftwareImageViaURLRequest).
		SetResult(&TaskIdResult{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TaskIdResult)
	return result, response, err

}

// TriggerSoftwareImageActivationQueryParams defines the query parameters for this request
type TriggerSoftwareImageActivationQueryParams struct {
	ScheduleValidate boolean `url:"scheduleValidate,omitempty"` // scheduleValidate, validates data before schedule (Optional)
}

// TriggerSoftwareImageActivation triggerSoftwareImageActivation
/* Activates a software image on a given device. Software image must be present in the device flash
@param Client-Type Client-type (Optional)
@param Client-Url Client-url (Optional)
@param scheduleValidate scheduleValidate, validates data before schedule (Optional)
*/
func (s *SoftwareImageManagementSWIMService) TriggerSoftwareImageActivation(triggerSoftwareImageActivationQueryParams *TriggerSoftwareImageActivationQueryParams, triggerSoftwareImageActivationRequest *TriggerSoftwareImageActivationRequest) (*TaskIdResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/image/activation/device"

	queryString, _ := query.Values(triggerSoftwareImageActivationQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetBody(triggerSoftwareImageActivationRequest).
		SetResult(&TaskIdResult{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TaskIdResult)
	return result, response, err

}

// TriggerSoftwareImageDistribution triggerSoftwareImageDistribution
/* Distributes a software image on a given device. Software image must be imported successfully into DNA Center before it can be distributed
 */
func (s *SoftwareImageManagementSWIMService) TriggerSoftwareImageDistribution(triggerSoftwareImageDistributionRequest *TriggerSoftwareImageDistributionRequest) (*TaskIdResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/image/distribution"

	response, err := RestyClient.R().
		SetBody(triggerSoftwareImageDistributionRequest).
		SetResult(&TaskIdResult{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TaskIdResult)
	return result, response, err

}