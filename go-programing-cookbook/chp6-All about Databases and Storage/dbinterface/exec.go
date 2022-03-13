package dbinterface

func Exec(db DB) error {
	// drop the table after the testing.
	defer db.Exec("DROP TABLE example")
	if err := Create(db); err != nil {
		return err
	}
	if err := Query(db); err != nil {
		return err
	}
	return nil
}
