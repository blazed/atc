BEGIN;

  ALTER TABLE independent_build_inputs RENAME job_combination_id TO job_id;

  ALTER TABLE independent_build_inputs DROP CONSTRAINT independent_build_inputs_unique_job_combination_id_input_name;

  ALTER TABLE independent_build_inputs ADD CONSTRAINT independent_build_inputs_unique_job_id_input_name UNIQUE (job_id, input_name);

  ALTER TABLE independent_build_inputs DROP CONSTRAINT independent_build_inputs_job_combination_id_fkey;

  ALTER TABLE independent_build_inputs ADD CONSTRAINT independent_build_inputs_job_id_fkey FOREIGN KEY (job_id) REFERENCES jobs(id) ON DELETE CASCADE;

  ALTER INDEX independent_build_inputs_job_combination_id RENAME TO independent_build_inputs_job_id;

COMMIT;
