package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Useractivity
type Useractivity struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The ID of the user
	Id *string `json:"id,omitempty"`

	// RoutingStatus - The current routing status of the user
	RoutingStatus *Useractivityroutingstatus `json:"routingStatus,omitempty"`

	// Presence - The current system presence of the user
	Presence *Useractivityadherencepresence `json:"presence,omitempty"`

	// OutOfOffice - The current out of office state of the user
	OutOfOffice *Useractivityoutofoffice `json:"outOfOffice,omitempty"`

	// ActiveQueueIds - The IDs of the queues for which the user is active
	ActiveQueueIds *[]string `json:"activeQueueIds,omitempty"`

	// DateActiveQueuesChanged - The date the activeQueueIds list was last modified. For reference only - subject to eventual consistency. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateActiveQueuesChanged *time.Time `json:"dateActiveQueuesChanged,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Useractivity) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Useractivity) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateActiveQueuesChanged", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Useractivity
	
	DateActiveQueuesChanged := new(string)
	if o.DateActiveQueuesChanged != nil {
		
		*DateActiveQueuesChanged = timeutil.Strftime(o.DateActiveQueuesChanged, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateActiveQueuesChanged = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		RoutingStatus *Useractivityroutingstatus `json:"routingStatus,omitempty"`
		
		Presence *Useractivityadherencepresence `json:"presence,omitempty"`
		
		OutOfOffice *Useractivityoutofoffice `json:"outOfOffice,omitempty"`
		
		ActiveQueueIds *[]string `json:"activeQueueIds,omitempty"`
		
		DateActiveQueuesChanged *string `json:"dateActiveQueuesChanged,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		RoutingStatus: o.RoutingStatus,
		
		Presence: o.Presence,
		
		OutOfOffice: o.OutOfOffice,
		
		ActiveQueueIds: o.ActiveQueueIds,
		
		DateActiveQueuesChanged: DateActiveQueuesChanged,
		Alias:    (Alias)(o),
	})
}

func (o *Useractivity) UnmarshalJSON(b []byte) error {
	var UseractivityMap map[string]interface{}
	err := json.Unmarshal(b, &UseractivityMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UseractivityMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if RoutingStatus, ok := UseractivityMap["routingStatus"].(map[string]interface{}); ok {
		RoutingStatusString, _ := json.Marshal(RoutingStatus)
		json.Unmarshal(RoutingStatusString, &o.RoutingStatus)
	}
	
	if Presence, ok := UseractivityMap["presence"].(map[string]interface{}); ok {
		PresenceString, _ := json.Marshal(Presence)
		json.Unmarshal(PresenceString, &o.Presence)
	}
	
	if OutOfOffice, ok := UseractivityMap["outOfOffice"].(map[string]interface{}); ok {
		OutOfOfficeString, _ := json.Marshal(OutOfOffice)
		json.Unmarshal(OutOfOfficeString, &o.OutOfOffice)
	}
	
	if ActiveQueueIds, ok := UseractivityMap["activeQueueIds"].([]interface{}); ok {
		ActiveQueueIdsString, _ := json.Marshal(ActiveQueueIds)
		json.Unmarshal(ActiveQueueIdsString, &o.ActiveQueueIds)
	}
	
	if dateActiveQueuesChangedString, ok := UseractivityMap["dateActiveQueuesChanged"].(string); ok {
		DateActiveQueuesChanged, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateActiveQueuesChangedString)
		o.DateActiveQueuesChanged = &DateActiveQueuesChanged
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Useractivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
