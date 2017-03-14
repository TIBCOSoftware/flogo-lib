package activity

import (
	"sync"
)

var (
	activitiesMu    sync.Mutex
	activities      = make(map[string]Activity)
	activityFactory = make(map[string]ActivityFactory)
)

// Register registers the specified activity
func Register(activity Activity) {
	activitiesMu.Lock()
	defer activitiesMu.Unlock()

	if activity == nil {
		panic("activity.Register: activity is nil")
	}

	id := activity.Metadata().ID

	if _, dup := activities[id]; dup {
		panic("activity.Register: activity already registered " + id)
	}

	// copy on write to avoid synchronization on access
	newActivities := make(map[string]Activity, len(activities))

	for k, v := range activities {
		newActivities[k] = v
	}

	newActivities[id] = activity
	activities = newActivities
}

func RegisterInstance(id string, activity Activity) {
	activitiesMu.Lock()
	defer activitiesMu.Unlock()

	if activity == nil {
		panic("activity.Register: activity is nil")
	}

	if _, dup := activities[id]; dup {
		panic("activity.Register: activity already registered " + id)
	}

	// copy on write to avoid synchronization on access
	newActivities := make(map[string]Activity, len(activities))

	for k, v := range activities {
		newActivities[k] = v
	}

	newActivities[id] = activity
	activities = newActivities
}

// Activities gets all the registered activities
func Activities() []Activity {

	var curActivities = activities

	list := make([]Activity, 0, len(curActivities))

	for _, value := range curActivities {
		list = append(list, value)
	}

	return list
}

// Get gets specified activity
func Get(id string) Activity {
	//var curActivities = activities
	return activities[id]
}

func RegisterFactory(id string, factory ActivityFactory) {
	activitiesMu.Lock()
	defer activitiesMu.Unlock()

	if factory == nil {
		panic("activity.RegisterFcatory: factory is nil")
	}

	if _, dup := activityFactory[id]; dup {
		panic("activity.RegisterFcatory: factory already registered " + id)
	}

	// copy on write to avoid synchronization on access
	newFactories := make(map[string]ActivityFactory, len(activityFactory))

	for k, v := range activityFactory {
		newFactories[k] = v
	}

	newFactories[id] = factory
	activityFactory = newFactories
}

// Activities gets all the registered activities
func ActivityFactories() []ActivityFactory {

	var curActivities = activityFactory

	list := make([]ActivityFactory, 0, len(curActivities))

	for _, value := range curActivities {
		list = append(list, value)
	}

	return list
}

// Get gets specified activity
func GetFactory(id string) ActivityFactory {
	//var curActivities = activities
	return activityFactory[id]
}
