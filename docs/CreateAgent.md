# CreateAgent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantName** | **string** |  | 
**ToolName** | **string** |  | 
**KnowledgebaseName** | Pointer to **string** |  | [optional] [default to ""]
**Description** | Pointer to **string** |  | [optional] [default to ""]
**Service** | Pointer to **string** |  | [optional] [default to ""]
**ChatDescription** | Pointer to **string** |  | [optional] [default to ""]
**CommunityDescription** | Pointer to **string** |  | [optional] [default to ""]
**SurveyDescription** | Pointer to **string** |  | [optional] [default to ""]
**ServiceDescription** | Pointer to **string** |  | [optional] [default to ""]
**Pdf** | Pointer to **string** |  | [optional] 
**Csv** | Pointer to **string** |  | [optional] 
**Excel** | Pointer to **string** |  | [optional] 
**Epub** | Pointer to **string** |  | [optional] 
**ImageFile** | Pointer to **string** |  | [optional] 
**WordDocument** | Pointer to **string** |  | [optional] 
**TxtFile** | Pointer to **string** |  | [optional] 
**Powerpoint** | Pointer to **string** |  | [optional] 
**GoogleDriveLink** | Pointer to **string** |  | [optional] 
**YoutubeLink** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**WikipediaQuery** | Pointer to **string** |  | [optional] 
**AuthToken** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateAgent

`func NewCreateAgent(tenantName string, toolName string, ) *CreateAgent`

NewCreateAgent instantiates a new CreateAgent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAgentWithDefaults

`func NewCreateAgentWithDefaults() *CreateAgent`

NewCreateAgentWithDefaults instantiates a new CreateAgent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantName

`func (o *CreateAgent) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *CreateAgent) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *CreateAgent) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.


### GetToolName

`func (o *CreateAgent) GetToolName() string`

GetToolName returns the ToolName field if non-nil, zero value otherwise.

### GetToolNameOk

`func (o *CreateAgent) GetToolNameOk() (*string, bool)`

GetToolNameOk returns a tuple with the ToolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolName

`func (o *CreateAgent) SetToolName(v string)`

SetToolName sets ToolName field to given value.


### GetKnowledgebaseName

`func (o *CreateAgent) GetKnowledgebaseName() string`

GetKnowledgebaseName returns the KnowledgebaseName field if non-nil, zero value otherwise.

### GetKnowledgebaseNameOk

`func (o *CreateAgent) GetKnowledgebaseNameOk() (*string, bool)`

GetKnowledgebaseNameOk returns a tuple with the KnowledgebaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgebaseName

`func (o *CreateAgent) SetKnowledgebaseName(v string)`

SetKnowledgebaseName sets KnowledgebaseName field to given value.

### HasKnowledgebaseName

`func (o *CreateAgent) HasKnowledgebaseName() bool`

HasKnowledgebaseName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateAgent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAgent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAgent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAgent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetService

`func (o *CreateAgent) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CreateAgent) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CreateAgent) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *CreateAgent) HasService() bool`

HasService returns a boolean if a field has been set.

### GetChatDescription

`func (o *CreateAgent) GetChatDescription() string`

GetChatDescription returns the ChatDescription field if non-nil, zero value otherwise.

### GetChatDescriptionOk

`func (o *CreateAgent) GetChatDescriptionOk() (*string, bool)`

GetChatDescriptionOk returns a tuple with the ChatDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatDescription

`func (o *CreateAgent) SetChatDescription(v string)`

SetChatDescription sets ChatDescription field to given value.

### HasChatDescription

`func (o *CreateAgent) HasChatDescription() bool`

HasChatDescription returns a boolean if a field has been set.

### GetCommunityDescription

`func (o *CreateAgent) GetCommunityDescription() string`

GetCommunityDescription returns the CommunityDescription field if non-nil, zero value otherwise.

### GetCommunityDescriptionOk

`func (o *CreateAgent) GetCommunityDescriptionOk() (*string, bool)`

GetCommunityDescriptionOk returns a tuple with the CommunityDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityDescription

`func (o *CreateAgent) SetCommunityDescription(v string)`

SetCommunityDescription sets CommunityDescription field to given value.

### HasCommunityDescription

`func (o *CreateAgent) HasCommunityDescription() bool`

HasCommunityDescription returns a boolean if a field has been set.

### GetSurveyDescription

`func (o *CreateAgent) GetSurveyDescription() string`

GetSurveyDescription returns the SurveyDescription field if non-nil, zero value otherwise.

### GetSurveyDescriptionOk

`func (o *CreateAgent) GetSurveyDescriptionOk() (*string, bool)`

GetSurveyDescriptionOk returns a tuple with the SurveyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyDescription

`func (o *CreateAgent) SetSurveyDescription(v string)`

SetSurveyDescription sets SurveyDescription field to given value.

### HasSurveyDescription

`func (o *CreateAgent) HasSurveyDescription() bool`

HasSurveyDescription returns a boolean if a field has been set.

### GetServiceDescription

`func (o *CreateAgent) GetServiceDescription() string`

GetServiceDescription returns the ServiceDescription field if non-nil, zero value otherwise.

### GetServiceDescriptionOk

`func (o *CreateAgent) GetServiceDescriptionOk() (*string, bool)`

GetServiceDescriptionOk returns a tuple with the ServiceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDescription

`func (o *CreateAgent) SetServiceDescription(v string)`

SetServiceDescription sets ServiceDescription field to given value.

### HasServiceDescription

`func (o *CreateAgent) HasServiceDescription() bool`

HasServiceDescription returns a boolean if a field has been set.

### GetPdf

`func (o *CreateAgent) GetPdf() string`

GetPdf returns the Pdf field if non-nil, zero value otherwise.

### GetPdfOk

`func (o *CreateAgent) GetPdfOk() (*string, bool)`

GetPdfOk returns a tuple with the Pdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdf

`func (o *CreateAgent) SetPdf(v string)`

SetPdf sets Pdf field to given value.

### HasPdf

`func (o *CreateAgent) HasPdf() bool`

HasPdf returns a boolean if a field has been set.

### GetCsv

`func (o *CreateAgent) GetCsv() string`

GetCsv returns the Csv field if non-nil, zero value otherwise.

### GetCsvOk

`func (o *CreateAgent) GetCsvOk() (*string, bool)`

GetCsvOk returns a tuple with the Csv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsv

`func (o *CreateAgent) SetCsv(v string)`

SetCsv sets Csv field to given value.

### HasCsv

`func (o *CreateAgent) HasCsv() bool`

HasCsv returns a boolean if a field has been set.

### GetExcel

`func (o *CreateAgent) GetExcel() string`

GetExcel returns the Excel field if non-nil, zero value otherwise.

### GetExcelOk

`func (o *CreateAgent) GetExcelOk() (*string, bool)`

GetExcelOk returns a tuple with the Excel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcel

`func (o *CreateAgent) SetExcel(v string)`

SetExcel sets Excel field to given value.

### HasExcel

`func (o *CreateAgent) HasExcel() bool`

HasExcel returns a boolean if a field has been set.

### GetEpub

`func (o *CreateAgent) GetEpub() string`

GetEpub returns the Epub field if non-nil, zero value otherwise.

### GetEpubOk

`func (o *CreateAgent) GetEpubOk() (*string, bool)`

GetEpubOk returns a tuple with the Epub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpub

`func (o *CreateAgent) SetEpub(v string)`

SetEpub sets Epub field to given value.

### HasEpub

`func (o *CreateAgent) HasEpub() bool`

HasEpub returns a boolean if a field has been set.

### GetImageFile

`func (o *CreateAgent) GetImageFile() string`

GetImageFile returns the ImageFile field if non-nil, zero value otherwise.

### GetImageFileOk

`func (o *CreateAgent) GetImageFileOk() (*string, bool)`

GetImageFileOk returns a tuple with the ImageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFile

`func (o *CreateAgent) SetImageFile(v string)`

SetImageFile sets ImageFile field to given value.

### HasImageFile

`func (o *CreateAgent) HasImageFile() bool`

HasImageFile returns a boolean if a field has been set.

### GetWordDocument

`func (o *CreateAgent) GetWordDocument() string`

GetWordDocument returns the WordDocument field if non-nil, zero value otherwise.

### GetWordDocumentOk

`func (o *CreateAgent) GetWordDocumentOk() (*string, bool)`

GetWordDocumentOk returns a tuple with the WordDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWordDocument

`func (o *CreateAgent) SetWordDocument(v string)`

SetWordDocument sets WordDocument field to given value.

### HasWordDocument

`func (o *CreateAgent) HasWordDocument() bool`

HasWordDocument returns a boolean if a field has been set.

### GetTxtFile

`func (o *CreateAgent) GetTxtFile() string`

GetTxtFile returns the TxtFile field if non-nil, zero value otherwise.

### GetTxtFileOk

`func (o *CreateAgent) GetTxtFileOk() (*string, bool)`

GetTxtFileOk returns a tuple with the TxtFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxtFile

`func (o *CreateAgent) SetTxtFile(v string)`

SetTxtFile sets TxtFile field to given value.

### HasTxtFile

`func (o *CreateAgent) HasTxtFile() bool`

HasTxtFile returns a boolean if a field has been set.

### GetPowerpoint

`func (o *CreateAgent) GetPowerpoint() string`

GetPowerpoint returns the Powerpoint field if non-nil, zero value otherwise.

### GetPowerpointOk

`func (o *CreateAgent) GetPowerpointOk() (*string, bool)`

GetPowerpointOk returns a tuple with the Powerpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerpoint

`func (o *CreateAgent) SetPowerpoint(v string)`

SetPowerpoint sets Powerpoint field to given value.

### HasPowerpoint

`func (o *CreateAgent) HasPowerpoint() bool`

HasPowerpoint returns a boolean if a field has been set.

### GetGoogleDriveLink

`func (o *CreateAgent) GetGoogleDriveLink() string`

GetGoogleDriveLink returns the GoogleDriveLink field if non-nil, zero value otherwise.

### GetGoogleDriveLinkOk

`func (o *CreateAgent) GetGoogleDriveLinkOk() (*string, bool)`

GetGoogleDriveLinkOk returns a tuple with the GoogleDriveLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleDriveLink

`func (o *CreateAgent) SetGoogleDriveLink(v string)`

SetGoogleDriveLink sets GoogleDriveLink field to given value.

### HasGoogleDriveLink

`func (o *CreateAgent) HasGoogleDriveLink() bool`

HasGoogleDriveLink returns a boolean if a field has been set.

### GetYoutubeLink

`func (o *CreateAgent) GetYoutubeLink() string`

GetYoutubeLink returns the YoutubeLink field if non-nil, zero value otherwise.

### GetYoutubeLinkOk

`func (o *CreateAgent) GetYoutubeLinkOk() (*string, bool)`

GetYoutubeLinkOk returns a tuple with the YoutubeLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYoutubeLink

`func (o *CreateAgent) SetYoutubeLink(v string)`

SetYoutubeLink sets YoutubeLink field to given value.

### HasYoutubeLink

`func (o *CreateAgent) HasYoutubeLink() bool`

HasYoutubeLink returns a boolean if a field has been set.

### GetUrl

`func (o *CreateAgent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateAgent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateAgent) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreateAgent) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetWikipediaQuery

`func (o *CreateAgent) GetWikipediaQuery() string`

GetWikipediaQuery returns the WikipediaQuery field if non-nil, zero value otherwise.

### GetWikipediaQueryOk

`func (o *CreateAgent) GetWikipediaQueryOk() (*string, bool)`

GetWikipediaQueryOk returns a tuple with the WikipediaQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWikipediaQuery

`func (o *CreateAgent) SetWikipediaQuery(v string)`

SetWikipediaQuery sets WikipediaQuery field to given value.

### HasWikipediaQuery

`func (o *CreateAgent) HasWikipediaQuery() bool`

HasWikipediaQuery returns a boolean if a field has been set.

### GetAuthToken

`func (o *CreateAgent) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *CreateAgent) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *CreateAgent) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *CreateAgent) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


