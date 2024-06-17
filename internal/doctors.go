package internal

import "database/sql"

type DoctorModel struct {
	DB *sql.DB
}

func (d *DoctorModel) CreateDoctor(doctor *Doctor) error {
	stmt := `INSERT INTO doctors (name, surname, position, age, experience)
		VALUES($1, $2, $3, $4, $5)`

	_, err := d.DB.Exec(stmt, doctor.Name, doctor.Surname, doctor.Position, doctor.Age, doctor.Experience)
	if err != nil {
		return err
	}

	return nil
}

func (d *DoctorModel) GetDoctorsSortByAgeDesc() ([]*Doctor, error) {
	stmt := `SELECT name, surname, position, age, experience FROM doctors ORDER BY age DESC`

	rows, err := d.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	doctors := []*Doctor{}

	for rows.Next() {
		dt := &Doctor{}
		err := rows.Scan(&dt.Name, &dt.Surname, &dt.Position, &dt.Age, &dt.Experience)
		if err != nil {
			return nil, err
		}
		doctors = append(doctors, dt)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return doctors, nil
}

func (d *DoctorModel) GetDoctorsSortByAgeAsc() ([]*Doctor, error) {
	stmt := `SELECT name, surname, position, age, experience FROM doctors ORDER BY age ASC`

	rows, err := d.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	doctors := []*Doctor{}

	for rows.Next() {
		dt := &Doctor{}
		err := rows.Scan(&dt.Name, &dt.Surname, &dt.Position, &dt.Age, &dt.Experience)
		if err != nil {
			return nil, err
		}
		doctors = append(doctors, dt)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return doctors, nil
}

func (d *DoctorModel) GetDoctorByName(name string) (*Doctor, error) {
	stmt := `SELECT name, surname, position, age, experience FROM doctors WHERE name = $1`

	dt := &Doctor{}
	err := d.DB.QueryRow(stmt, name).Scan(&dt.Name, &dt.Surname, &dt.Position, &dt.Age, &dt.Experience)
	if err != nil {
		return nil, err
	}

	return dt, nil
}
func (d *DoctorModel) UpdateDoctor(doctor *Doctor) error {
	stmt := `UPDATE doctors SET name = $1, surname = $2, position = $3, age = $4, experience = $5 WHERE name = $1`

	_, err := d.DB.Exec(stmt, doctor.Name, doctor.Surname, doctor.Position, doctor.Age, doctor.Experience)
	if err != nil {
		return err
	}

	return nil
}

func (d *DoctorModel) DeleteDoctor(name string) error {
	stmt := `DELETE FROM doctors WHERE name = $1`

	_, err := d.DB.Exec(stmt, name)
	if err != nil {
		return err
	}

	return nil
}
