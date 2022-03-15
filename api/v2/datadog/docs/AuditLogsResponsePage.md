# AuditLogsResponsePage

## Properties

| Name      | Type                  | Description                                                                                                                                         | Notes      |
| --------- | --------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| **After** | Pointer to **string** | The cursor to use to get the next results, if any. To make the next request, use the same parameters with the addition of &#x60;page[cursor]&#x60;. | [optional] |

## Methods

### NewAuditLogsResponsePage

`func NewAuditLogsResponsePage() *AuditLogsResponsePage`

NewAuditLogsResponsePage instantiates a new AuditLogsResponsePage object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAuditLogsResponsePageWithDefaults

`func NewAuditLogsResponsePageWithDefaults() *AuditLogsResponsePage`

NewAuditLogsResponsePageWithDefaults instantiates a new AuditLogsResponsePage object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAfter

`func (o *AuditLogsResponsePage) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *AuditLogsResponsePage) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *AuditLogsResponsePage) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *AuditLogsResponsePage) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)