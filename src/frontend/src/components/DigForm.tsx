
import { Button, TextField } from "@mui/material";
import { Form, Formik, Field } from "formik";
import * as React from "react";
import DigQueryModel from "../models/DigQueryModel";



interface Props {
    onSubmit: (query: DigQueryModel) => void;
}

export const DigForm: React.FC<Props> = ({ onSubmit }) => {
    return (
        <Formik
            initialValues={{ type: 1, domain: "" }}
            onSubmit={values => {
                // onSubmit(null);
            }}
        >
            {({ values }) => (
                <Form>
                    <div>
                        <Field
                            name="domain"
                            placeholder="domain"
                            component={TextField}
                        />
                    </div>

                    <Button type="submit">submit</Button>
                    <pre>{JSON.stringify(values, null, 2)}</pre>
                </Form>
            )}
        </Formik>
    );
};