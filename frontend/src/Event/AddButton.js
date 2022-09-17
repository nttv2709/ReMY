
import * as React from 'react';
import {Box, Modal, Typography} from "@mui/material";
import {TextField} from '@mui/material';

import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { DateTimePicker } from '@mui/x-date-pickers/DateTimePicker';

const style = {
    position: 'absolute',
    top: '50%',
    left: '50%',
    transform: 'translate(-50%, -50%)',
    width: 400,
    bgcolor: 'background.paper',
    border: '2px solid #000',
    boxShadow: 24,
    p: 4,
};

function AddButton() {

    return (
                <Modal
                    aria-labelledby="modal-modal-title"
                    aria-describedby="modal-modal-description"
                    open = "true">
                    <Box sx={style}>
                        <Typography id="modal-modal-title" variant="h6" component="h2">
                            Add event
                        </Typography>
                        <TextField id="standard-basic" label="Add the title here" variant="standard" />
                        <LocalizationProvider dateAdapter={AdapterDayjs}>
                            <Stack spacing={3}>
                                <DateTimePicker
                                    label="Date&Time picker"
                                    value={value}
                                    onChange={handleChange}
                                    renderInput={(params) => <TextField {...params} />}
                                />
                            </Stack>
                        </LocalizationProvider>
                    </Box>
                </Modal>
    );
}

export default AddButton;