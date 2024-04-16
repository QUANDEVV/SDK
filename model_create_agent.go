/*
proxima-core-engine

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CreateAgent struct for CreateAgent
type CreateAgent struct {
	TenantName string `json:"tenant_name"`
	ToolName string `json:"tool_name"`
	KnowledgebaseName *string `json:"knowledgebase_name,omitempty"`
	Description *string `json:"description,omitempty"`
	Service *string `json:"service,omitempty"`
	ChatDescription *string `json:"chat_description,omitempty"`
	CommunityDescription *string `json:"community_description,omitempty"`
	SurveyDescription *string `json:"survey_description,omitempty"`
	ServiceDescription *string `json:"service_description,omitempty"`
	Pdf *string `json:"pdf,omitempty"`
	Csv *string `json:"csv,omitempty"`
	Excel *string `json:"excel,omitempty"`
	Epub *string `json:"epub,omitempty"`
	ImageFile *string `json:"image_file,omitempty"`
	WordDocument *string `json:"word_document,omitempty"`
	TxtFile *string `json:"txt_file,omitempty"`
	Powerpoint *string `json:"powerpoint,omitempty"`
	GoogleDriveLink *string `json:"google_drive_link,omitempty"`
	YoutubeLink *string `json:"youtube_link,omitempty"`
	Url *string `json:"url,omitempty"`
	WikipediaQuery *string `json:"wikipedia_query,omitempty"`
	AuthToken *string `json:"auth_token,omitempty"`
}

// NewCreateAgent instantiates a new CreateAgent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAgent(tenantName string, toolName string) *CreateAgent {
	this := CreateAgent{}
	this.TenantName = tenantName
	this.ToolName = toolName
	var knowledgebaseName string = ""
	this.KnowledgebaseName = &knowledgebaseName
	var description string = ""
	this.Description = &description
	var service string = ""
	this.Service = &service
	var chatDescription string = ""
	this.ChatDescription = &chatDescription
	var communityDescription string = ""
	this.CommunityDescription = &communityDescription
	var surveyDescription string = ""
	this.SurveyDescription = &surveyDescription
	var serviceDescription string = ""
	this.ServiceDescription = &serviceDescription
	return &this
}

// NewCreateAgentWithDefaults instantiates a new CreateAgent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAgentWithDefaults() *CreateAgent {
	this := CreateAgent{}
	var knowledgebaseName string = ""
	this.KnowledgebaseName = &knowledgebaseName
	var description string = ""
	this.Description = &description
	var service string = ""
	this.Service = &service
	var chatDescription string = ""
	this.ChatDescription = &chatDescription
	var communityDescription string = ""
	this.CommunityDescription = &communityDescription
	var surveyDescription string = ""
	this.SurveyDescription = &surveyDescription
	var serviceDescription string = ""
	this.ServiceDescription = &serviceDescription
	return &this
}

// GetTenantName returns the TenantName field value
func (o *CreateAgent) GetTenantName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantName
}

// GetTenantNameOk returns a tuple with the TenantName field value
// and a boolean to check if the value has been set.
func (o *CreateAgent) GetTenantNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TenantName, true
}

// SetTenantName sets field value
func (o *CreateAgent) SetTenantName(v string) {
	o.TenantName = v
}

// GetToolName returns the ToolName field value
func (o *CreateAgent) GetToolName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToolName
}

// GetToolNameOk returns a tuple with the ToolName field value
// and a boolean to check if the value has been set.
func (o *CreateAgent) GetToolNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ToolName, true
}

// SetToolName sets field value
func (o *CreateAgent) SetToolName(v string) {
	o.ToolName = v
}

// GetKnowledgebaseName returns the KnowledgebaseName field value if set, zero value otherwise.
func (o *CreateAgent) GetKnowledgebaseName() string {
	if o == nil || o.KnowledgebaseName == nil {
		var ret string
		return ret
	}
	return *o.KnowledgebaseName
}

// GetKnowledgebaseNameOk returns a tuple with the KnowledgebaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgent) GetKnowledgebaseNameOk() (*string, bool) {
	if o == nil || o.KnowledgebaseName == nil {
		return nil, false
	}
	return o.KnowledgebaseName, true
}

// HasKnowledgebaseName returns a boolean if a field has been set.
func (o *CreateAgent) HasKnowledgebaseName() bool {
	if o != nil && o.KnowledgebaseName != nil {
		return true
	}

	return false
}

// SetKnowledgebaseName gets a reference to the given string and assigns it to the KnowledgebaseName field.
func (o *CreateAgent) SetKnowledgebaseName(v string) {
	o.KnowledgebaseName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateAgent) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgent) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateAgent) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateAgent) SetDescription(v string) {
	o.Description = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *CreateAgent) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgent) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *CreateAgent) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *CreateAgent) SetService(v string) {
	o.Service = &v
}

// GetChatDescription returns the ChatDescription field value if set, zero value otherwise.
func (o *CreateAgent) GetChatDescription() string {
	if o == nil || o.ChatDescription == nil {
		var ret string
		return ret
	}
	return *o.ChatDescription
}

// GetChatDescriptionOk returns a tuple with the ChatDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgent) GetChatDescriptionOk() (*string, bool) {
	if o == nil || o.ChatDescription == nil {
		return nil, false
	}
	return o.ChatDescription, true
}

// HasChatDescription returns a boolean if a field has been set.
func (o *CreateAgent) HasChatDescription() bool {
	if o != nil && o.ChatDescription != nil {
		return true
	}

	return false
}

// SetChatDescription gets a reference to the given string and assigns it to the ChatDescription field.
func (o *CreateAgent) SetChatDescription(v string) {
	o.ChatDescription = &v
}

// GetCommunityDescription returns the CommunityDescription field value if set, zero value otherwise.
func (o *CreateAgent) GetCommunityDescription() string {
	if o == nil || o.CommunityDescription == nil {
		var ret string
		return ret
	}
	return *o.CommunityDescription
}

// GetCommunityDescriptionOk returns a tuple with the CommunityDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgent) GetCommunityDescriptionOk() (*string, bool) {
	if o == nil || o.CommunityDescription == nil {
		return nil, false
	}
	return o.CommunityDescription, true
}

// HasCommunityDescription returns a boolean if a field has been set.
func (o *CreateAgent) HasCommunityDescription() bool {
	if o != nil && o.CommunityDescription != nil {
		return true
	}

	return false
}

// SetCommunityDescription gets a reference to the given string and assigns it to the CommunityDescription field.
func (o *CreateAgent) SetCommunityDescription(v string) {
	o.CommunityDescription = &v
}

// GetSurveyDescription returns the SurveyDescription field value if set, zero value otherwise.
func (o *CreateAgent) GetSurveyDescription() string {
	if o == nil || o.SurveyDescription == nil {
		var ret string
		return ret
	}
	return *o.SurveyDescription
}

// GetSurveyDescriptionOk returns a tuple with the SurveyDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgent) GetSurveyDescriptionOk() (*string, bool) {
	if o == nil || o.SurveyDescription == nil {
		return nil, false
	}
	return o.SurveyDescription, true
}

// HasSurveyDescription returns a boolean if a field has been set.
func (o *CreateAgent) HasSurveyDescription() bool {
	if o != nil && o.SurveyDescription != nil {
		return true
	}

	return false
}

// SetSurveyDescription gets a reference to the given string and assigns it to the SurveyDescription field.
func (o *CreateAgent) SetSurveyDescription(v string) {
	o.SurveyDescription = &v
}

// GetServiceDescription returns the ServiceDescription field value if set, zero value otherwise.
func (o *CreateAgent) GetServiceDescription() string {
	if o == nil || o.ServiceDescription == nil {
		var ret string
		return ret
	}
	return *o.ServiceDescription
}

// GetServiceDescriptionOk returns a tuple with the ServiceDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgent) GetServiceDescriptionOk() (*string, bool) {
	if o == nil || o.ServiceDescription == nil {
		return nil, false
	}
	return o.ServiceDescription, true
}

// HasServiceDescription returns a boolean if a field has been set.
func (o *CreateAgent) HasServiceDescription() bool {
	if o != nil && o.ServiceDescription != nil {
		return true
	}

	return false
}

// SetServiceDescription gets a reference to the given string and assigns it to the ServiceDescription field.
func (o *CreateAgent) SetServiceDescription(v string) {
	o.ServiceDescription = &v
}

// GetPdf returns the Pdf field value if set, zero value otherwise.
func (o *CreateAgent) GetPdf() string {
	if o == nil || o.Pdf == nil {
		var ret string
		return ret
	}
	return *o.Pdf
}

// GetPdfOk returns a tuple with the Pdf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgent) GetPdfOk() (*string, bool) {
	if o == nil || o.Pdf == nil {
		return nil, false
	}
	return o.Pdf, true
}

// HasPdf returns a boolean if a field has been set.
func (o *CreateAgent) HasPdf() bool {
	if o != nil && o.Pdf != nil {
		return true
	}

	return false
}

// SetPdf gets a reference to the given string and assigns it to the Pdf field.
func (o *CreateAgent) SetPdf(v string) {
	o.Pdf = &v
}

// GetCsv returns the Csv field value if set, zero value otherwise.
func (o *CreateAgent) GetCsv() string {
	if o == nil || o.Csv == nil {
		var ret string
		return ret
	}
	return *o.Csv
}

// GetCsvOk returns a tuple with the Csv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgent) GetCsvOk() (*string, bool) {
	if o == nil || o.Csv == nil {
		return nil, false
	}
	return o.Csv, true
}

// HasCsv returns a boolean if a field has been set.
func (o *CreateAgent) HasCsv() bool {
	if o != nil && o.Csv != nil {
		return true
	}

	return false
}

// SetCsv gets a reference to the given string and assigns it to the Csv field.
func (o *CreateAgent) SetCsv(v string) {
	o.Csv = &v
}

// GetExcel returns the Excel field value if set, zero value otherwise.
func (o *CreateAgent) GetExcel() string {
	if o == nil || o.Excel == nil {
		var ret string
		return ret
	}
	return *o.Excel
}

// GetExcelOk returns a tuple with the Excel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgent) GetExcelOk() (*string, bool) {
	if o == nil || o.Excel == nil {
		return nil, false
	}
	return o.Excel, true
}

// HasExcel returns a boolean if a field has been set.
func (o *CreateAgent) HasExcel() bool {
	if o != nil && o.Excel != nil {
		return true
	}

	return false
}

// SetExcel gets a reference to the given string and assigns it to the Excel field.
func (o *CreateAgent) SetExcel(v string) {
	o.Excel = &v
}

// GetEpub returns the Epub field value if set, zero value otherwise.
func (o *CreateAgent) GetEpub() string {
	if o == nil || o.Epub == nil {
		var ret string
		return ret
	}
	return *o.Epub
}

// GetEpubOk returns a tuple with the Epub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgent) GetEpubOk() (*string, bool) {
	if o == nil || o.Epub == nil {
		return nil, false
	}
	return o.Epub, true
}

// HasEpub returns a boolean if a field has been set.
func (o *CreateAgent) HasEpub() bool {
	if o != nil && o.Epub != nil {
		return true
	}

	return false
}

// SetEpub gets a reference to the given string and assigns it to the Epub field.
func (o *CreateAgent) SetEpub(v string) {
	o.Epub = &v
}

// GetImageFile returns the ImageFile field value if set, zero value otherwise.
func (o *CreateAgent) GetImageFile() string {
	if o == nil || o.ImageFile == nil {
		var ret string
		return ret
	}
	return *o.ImageFile
}

// GetImageFileOk returns a tuple with the ImageFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgent) GetImageFileOk() (*string, bool) {
	if o == nil || o.ImageFile == nil {
		return nil, false
	}
	return o.ImageFile, true
}

// HasImageFile returns a boolean if a field has been set.
func (o *CreateAgent) HasImageFile() bool {
	if o != nil && o.ImageFile != nil {
		return true
	}

	return false
}

// SetImageFile gets a reference to the given string and assigns it to the ImageFile field.
func (o *CreateAgent) SetImageFile(v string) {
	o.ImageFile = &v
}

// GetWordDocument returns the WordDocument field value if set, zero value otherwise.
func (o *CreateAgent) GetWordDocument() string {
	if o == nil || o.WordDocument == nil {
		var ret string
		return ret
	}
	return *o.WordDocument
}

// GetWordDocumentOk returns a tuple with the WordDocument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgent) GetWordDocumentOk() (*string, bool) {
	if o == nil || o.WordDocument == nil {
		return nil, false
	}
	return o.WordDocument, true
}

// HasWordDocument returns a boolean if a field has been set.
func (o *CreateAgent) HasWordDocument() bool {
	if o != nil && o.WordDocument != nil {
		return true
	}

	return false
}

// SetWordDocument gets a reference to the given string and assigns it to the WordDocument field.
func (o *CreateAgent) SetWordDocument(v string) {
	o.WordDocument = &v
}

// GetTxtFile returns the TxtFile field value if set, zero value otherwise.
func (o *CreateAgent) GetTxtFile() string {
	if o == nil || o.TxtFile == nil {
		var ret string
		return ret
	}
	return *o.TxtFile
}

// GetTxtFileOk returns a tuple with the TxtFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgent) GetTxtFileOk() (*string, bool) {
	if o == nil || o.TxtFile == nil {
		return nil, false
	}
	return o.TxtFile, true
}

// HasTxtFile returns a boolean if a field has been set.
func (o *CreateAgent) HasTxtFile() bool {
	if o != nil && o.TxtFile != nil {
		return true
	}

	return false
}

// SetTxtFile gets a reference to the given string and assigns it to the TxtFile field.
func (o *CreateAgent) SetTxtFile(v string) {
	o.TxtFile = &v
}

// GetPowerpoint returns the Powerpoint field value if set, zero value otherwise.
func (o *CreateAgent) GetPowerpoint() string {
	if o == nil || o.Powerpoint == nil {
		var ret string
		return ret
	}
	return *o.Powerpoint
}

// GetPowerpointOk returns a tuple with the Powerpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgent) GetPowerpointOk() (*string, bool) {
	if o == nil || o.Powerpoint == nil {
		return nil, false
	}
	return o.Powerpoint, true
}

// HasPowerpoint returns a boolean if a field has been set.
func (o *CreateAgent) HasPowerpoint() bool {
	if o != nil && o.Powerpoint != nil {
		return true
	}

	return false
}

// SetPowerpoint gets a reference to the given string and assigns it to the Powerpoint field.
func (o *CreateAgent) SetPowerpoint(v string) {
	o.Powerpoint = &v
}

// GetGoogleDriveLink returns the GoogleDriveLink field value if set, zero value otherwise.
func (o *CreateAgent) GetGoogleDriveLink() string {
	if o == nil || o.GoogleDriveLink == nil {
		var ret string
		return ret
	}
	return *o.GoogleDriveLink
}

// GetGoogleDriveLinkOk returns a tuple with the GoogleDriveLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgent) GetGoogleDriveLinkOk() (*string, bool) {
	if o == nil || o.GoogleDriveLink == nil {
		return nil, false
	}
	return o.GoogleDriveLink, true
}

// HasGoogleDriveLink returns a boolean if a field has been set.
func (o *CreateAgent) HasGoogleDriveLink() bool {
	if o != nil && o.GoogleDriveLink != nil {
		return true
	}

	return false
}

// SetGoogleDriveLink gets a reference to the given string and assigns it to the GoogleDriveLink field.
func (o *CreateAgent) SetGoogleDriveLink(v string) {
	o.GoogleDriveLink = &v
}

// GetYoutubeLink returns the YoutubeLink field value if set, zero value otherwise.
func (o *CreateAgent) GetYoutubeLink() string {
	if o == nil || o.YoutubeLink == nil {
		var ret string
		return ret
	}
	return *o.YoutubeLink
}

// GetYoutubeLinkOk returns a tuple with the YoutubeLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgent) GetYoutubeLinkOk() (*string, bool) {
	if o == nil || o.YoutubeLink == nil {
		return nil, false
	}
	return o.YoutubeLink, true
}

// HasYoutubeLink returns a boolean if a field has been set.
func (o *CreateAgent) HasYoutubeLink() bool {
	if o != nil && o.YoutubeLink != nil {
		return true
	}

	return false
}

// SetYoutubeLink gets a reference to the given string and assigns it to the YoutubeLink field.
func (o *CreateAgent) SetYoutubeLink(v string) {
	o.YoutubeLink = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CreateAgent) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgent) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CreateAgent) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CreateAgent) SetUrl(v string) {
	o.Url = &v
}

// GetWikipediaQuery returns the WikipediaQuery field value if set, zero value otherwise.
func (o *CreateAgent) GetWikipediaQuery() string {
	if o == nil || o.WikipediaQuery == nil {
		var ret string
		return ret
	}
	return *o.WikipediaQuery
}

// GetWikipediaQueryOk returns a tuple with the WikipediaQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgent) GetWikipediaQueryOk() (*string, bool) {
	if o == nil || o.WikipediaQuery == nil {
		return nil, false
	}
	return o.WikipediaQuery, true
}

// HasWikipediaQuery returns a boolean if a field has been set.
func (o *CreateAgent) HasWikipediaQuery() bool {
	if o != nil && o.WikipediaQuery != nil {
		return true
	}

	return false
}

// SetWikipediaQuery gets a reference to the given string and assigns it to the WikipediaQuery field.
func (o *CreateAgent) SetWikipediaQuery(v string) {
	o.WikipediaQuery = &v
}

// GetAuthToken returns the AuthToken field value if set, zero value otherwise.
func (o *CreateAgent) GetAuthToken() string {
	if o == nil || o.AuthToken == nil {
		var ret string
		return ret
	}
	return *o.AuthToken
}

// GetAuthTokenOk returns a tuple with the AuthToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgent) GetAuthTokenOk() (*string, bool) {
	if o == nil || o.AuthToken == nil {
		return nil, false
	}
	return o.AuthToken, true
}

// HasAuthToken returns a boolean if a field has been set.
func (o *CreateAgent) HasAuthToken() bool {
	if o != nil && o.AuthToken != nil {
		return true
	}

	return false
}

// SetAuthToken gets a reference to the given string and assigns it to the AuthToken field.
func (o *CreateAgent) SetAuthToken(v string) {
	o.AuthToken = &v
}

func (o CreateAgent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tenant_name"] = o.TenantName
	}
	if true {
		toSerialize["tool_name"] = o.ToolName
	}
	if o.KnowledgebaseName != nil {
		toSerialize["knowledgebase_name"] = o.KnowledgebaseName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.ChatDescription != nil {
		toSerialize["chat_description"] = o.ChatDescription
	}
	if o.CommunityDescription != nil {
		toSerialize["community_description"] = o.CommunityDescription
	}
	if o.SurveyDescription != nil {
		toSerialize["survey_description"] = o.SurveyDescription
	}
	if o.ServiceDescription != nil {
		toSerialize["service_description"] = o.ServiceDescription
	}
	if o.Pdf != nil {
		toSerialize["pdf"] = o.Pdf
	}
	if o.Csv != nil {
		toSerialize["csv"] = o.Csv
	}
	if o.Excel != nil {
		toSerialize["excel"] = o.Excel
	}
	if o.Epub != nil {
		toSerialize["epub"] = o.Epub
	}
	if o.ImageFile != nil {
		toSerialize["image_file"] = o.ImageFile
	}
	if o.WordDocument != nil {
		toSerialize["word_document"] = o.WordDocument
	}
	if o.TxtFile != nil {
		toSerialize["txt_file"] = o.TxtFile
	}
	if o.Powerpoint != nil {
		toSerialize["powerpoint"] = o.Powerpoint
	}
	if o.GoogleDriveLink != nil {
		toSerialize["google_drive_link"] = o.GoogleDriveLink
	}
	if o.YoutubeLink != nil {
		toSerialize["youtube_link"] = o.YoutubeLink
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.WikipediaQuery != nil {
		toSerialize["wikipedia_query"] = o.WikipediaQuery
	}
	if o.AuthToken != nil {
		toSerialize["auth_token"] = o.AuthToken
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAgent struct {
	value *CreateAgent
	isSet bool
}

func (v NullableCreateAgent) Get() *CreateAgent {
	return v.value
}

func (v *NullableCreateAgent) Set(val *CreateAgent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAgent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAgent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAgent(val *CreateAgent) *NullableCreateAgent {
	return &NullableCreateAgent{value: val, isSet: true}
}

func (v NullableCreateAgent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAgent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

