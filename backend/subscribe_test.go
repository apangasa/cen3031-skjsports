package main

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestInitialSubscription(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("skjsports.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	email := "test.initial@subscription.com"
	first := "Test"
	last := "Subscription"

	delete_query := "DELETE FROM Subscribers WHERE ID = '" + email + "'"

	db.Exec(delete_query)
	success := addSubscriber(email, first, last)

	if !success {
		db.Exec(delete_query)
		t.Fatalf(`addSubscriber("test.initial@subscription.com", "Test", "Subscription") affects 0 rows`)
	} else {
		result := retrieveSubscriber(email)
		db.Exec(delete_query)
		if result == nil {
			t.Fatalf(`addSubscriber("test.initial@subscription.com", "Test", "Subscription") results in retrieveSubscriber("test.initial@subscription.com") = %v`, nil)
		} else {
			if result.ID != email || result.FirstName != "Test" || result.LastName != "Subscription" {
				t.Fatalf(`addSubscriber("test.initial@subscription.com", "Test", "Subscription") results in retrieveSubscriber("test.initial@subscription.com" with incorrect attributes`)
			}
		}
	}
}

func TestResubscriptionWhileSubscribed(t *testing.T) {
	/*
		Resubscribing while already subscribed should fail, so new personal details should not be stored in DB.
	*/
	db, err := gorm.Open(sqlite.Open("skjsports.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	email := "test.re@subscription.com"
	first := "Test"
	last := "Subscription"

	delete_query := "DELETE FROM Subscribers WHERE ID = '" + email + "'"

	db.Exec(delete_query)
	addSubscriber(email, first, last)

	first = "New"
	last = "Info"

	success := addSubscriber(email, first, last)

	if success {
		db.Exec(delete_query)
		t.Fatalf(`addSubscriber("test.re@subscription.com", "New", "Info") succeeds when "test.re@subscription.com" is already in database`)
	} else {
		result := retrieveSubscriber(email)
		db.Exec(delete_query)
		if result == nil {
			t.Fatalf(`addSubscriber("test.re@subscription.com", "New", "Info") results in retrieveSubscriber("test.re@subscription.com") = %v`, nil)
		} else {
			if result.ID != email || result.FirstName != "Test" || result.LastName != "Subscription" {
				t.Fatalf(`addSubscriber("test.re@subscription.com", "New", "Info") results in retrieveSubscriber("test.re@subscription.com") with incorrect attributes`)
			}
		}
	}
}

func TestUnsubscription(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("skjsports.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	email := "test.un@subscription.com"
	first := "Test"
	last := "Subscription"

	delete_query := "DELETE FROM Subscribers WHERE ID = '" + email + "'"

	db.Exec(delete_query)
	addSubscriber(email, first, last)
	success := removeSubscriber(email)

	if !success {
		db.Exec(delete_query)
		t.Fatalf(`removeSubscriber("test.un@subscription.com", "Test", "Subscription") affects 0 rows`)
	} else {
		result := retrieveSubscriber(email)
		db.Exec(delete_query)
		if result != nil {
			t.Fatalf(`removeSubscriber("test.un@subscription.com") results in retrieveSubscriber("test.un@subscription.com") with results`)
		}
	}
}

func TestResubscriptionAfterUnsubscription(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("skjsports.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	email := "test.re@subscription.com"
	first := "Test"
	last := "Subscription"

	delete_query := "DELETE FROM Subscribers WHERE ID = '" + email + "'"

	db.Exec(delete_query)
	addSubscriber(email, first, last)
	removeSubscriber(email)
	success := addSubscriber(email, first, last)

	if !success {
		db.Exec(delete_query)
		t.Fatalf(`addSubscriber("test.re@subscription.com", "Test", "Subscription") affects 0 rows`)
	} else {
		result := retrieveSubscriber(email)
		db.Exec(delete_query)
		if result == nil {
			t.Fatalf(`addSubscriber("test.re@subscription.com", "Test", "Subscription") results in retrieveSubscriber("test.re@subscription.com") = %v`, nil)
		} else {
			if result.ID != email || result.FirstName != "Test" || result.LastName != "Subscription" {
				t.Fatalf(`addSubscriber("test.re@subscription.com", "Test", "Subscription") results in retrieveSubscriber("test.re@subscription.com") with incorrect attributes`)
			}
		}
	}
}

func TestResubscriptionWithDifferentDetails(t *testing.T) {
	/*
		This time personal details can change since the user is legitimately resubscribing
	*/
	db, err := gorm.Open(sqlite.Open("skjsports.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	email := "test.re@subscription.com"
	first := "Test"
	last := "Subscription"

	delete_query := "DELETE FROM Subscribers WHERE ID = '" + email + "'"

	db.Exec(delete_query)
	addSubscriber(email, first, last)
	removeSubscriber(email)

	first = "New"
	last = "Info"
	success := addSubscriber(email, first, last)

	if !success {
		db.Exec(delete_query)
		t.Fatalf(`addSubscriber("test.re@subscription.com", "New", "Info") affects 0 rows`)
	} else {
		result := retrieveSubscriber(email)
		db.Exec(delete_query)
		if result == nil {
			t.Fatalf(`addSubscriber("test.re@subscription.com", "New", "Info") results in retrieveSubscriber("test.re@subscription.com") = %v`, nil)
		} else {
			if result.ID != email || result.FirstName != "New" || result.LastName != "Info" {
				t.Fatalf(`addSubscriber("test.re@subscription.com", "New", "Info") results in retrieveSubscriber("test.re@subscription.com") with incorrect attributes`)
			}
		}
	}
}
