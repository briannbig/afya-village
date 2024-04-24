package dto

type (
	RequestCreateUser struct {
		Username    string `json:"username"`
		Email       string `json:"email"`
		Telephone   string `json:"telephone"`
		Password    string `json:"password"`
		DateOfBirth string `json:"dateOfBirth"`
	}

	RequestCreateMedicalRecord struct {
		PatientID       string   `json:"patientId"`
		DateTime        string   `json:"timestamp"`
		Diagnosis       string   `json:"diagnosis"`
		Treatment       string   `json:"treatment"`
		MedicId         string   `json:"medicId"`
		PrescriptionIds []string `json:"prescriptionIds"`
	}

	RequestCreatePatient struct {
		Name        string `json:"name"`
		DateOfBirth string `json:"dateOfBirth"`
		Gender      string `json:"gender"`
		Location    string `json:"location"`
	}

	RequestCreateAppointment struct {
		PatientID           string `json:"patientId"`
		MobileClinicID      string `json:"mobileClinicId"`
		AppointmentDateTime string `json:"appointmentTime"`
	}
)
