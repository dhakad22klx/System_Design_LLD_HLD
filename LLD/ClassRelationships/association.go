// /One object need to know about the existence of another object to perform its responsibilities

package main

import (
	"fmt"
)

// Room holds the room's number and floor
type Room struct {
	Number string
	Floor  int
}

// Doctor holds name, specialization, and a slice of appointment pointers
type Doctor struct {
	Name           string
	Specialization string
	Appointments   []*Appointment
}

// AddAppointment appends an appointment to the doctor's list
func (d *Doctor) AddAppointment(a *Appointment) {
	d.Appointments = append(d.Appointments, a)
}

// GetPatients returns a deduplicated slice of patients this doctor has seen
func (d *Doctor) GetPatients() []*Patient {
	seen := make(map[*Patient]bool)
	var result []*Patient
	for _, a := range d.Appointments {
		if !seen[a.Patient] {
			seen[a.Patient] = true
			result = append(result, a.Patient)
		}
	}
	return result
}

// Patient holds a name and a slice of appointment pointers
type Patient struct {
	Name         string
	Appointments []*Appointment
}

// AddAppointment appends an appointment to the patient's list
func (p *Patient) AddAppointment(a *Appointment) {
	p.Appointments = append(p.Appointments, a)
}

// GetDoctors returns a deduplicated slice of doctors this patient has seen
func (p *Patient) GetDoctors() []*Doctor {
	seen := make(map[*Doctor]bool)
	var result []*Doctor
	for _, a := range p.Appointments {
		if !seen[a.Doctor] {
			result = append(result, a.Doctor)
		}
	}
	return result
}

// Appointment is the associative struct linking Doctor, Patient, and Room
type Appointment struct {
	Doctor  *Doctor
	Patient *Patient
	Room    *Room
	Time    string
}

// NewAppointment creates an Appointment and registers it with both Doctor and Patient
func NewAppointment(doctor *Doctor, patient *Patient, room *Room, time string) *Appointment {
	a := &Appointment{
		Doctor:  doctor,
		Patient: patient,
		Room:    room,
		Time:    time,
	}
	doctor.AddAppointment(a)
	patient.AddAppointment(a)
	return a
}

func testAssociation() {
	drSmith := &Doctor{Name: "Dr. Smith", Specialization: "Cardiology"}
	drPatel := &Doctor{Name: "Dr. Patel", Specialization: "Neurology"}

	alice := &Patient{Name: "Alice"}
	bob := &Patient{Name: "Bob"}

	room101 := &Room{Number: "101", Floor: 1}
	room205 := &Room{Number: "205", Floor: 2}

	NewAppointment(drSmith, alice, room101, "9:00 AM")
	NewAppointment(drSmith, bob, room101, "10:00 AM")
	NewAppointment(drPatel, alice, room205, "2:00 PM")

	fmt.Printf("%s's patients:\n", drSmith.Name)
	for _, p := range drSmith.GetPatients() {
		fmt.Printf("  - %s\n", p.Name)
	}

	fmt.Printf("%s's doctors:\n", alice.Name)
	for _, d := range alice.GetDoctors() {
		fmt.Printf("  - %s (%s)\n", d.Name, d.Specialization)
	}

	fmt.Printf("%s's schedule:\n", drSmith.Name)
	for _, a := range drSmith.Appointments {
		fmt.Printf("  - %s with %s in Room %s\n", a.Time, a.Patient.Name, a.Room.Number)
	}
}
