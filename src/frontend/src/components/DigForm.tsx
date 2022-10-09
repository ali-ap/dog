import { ContentPasteSearchRounded } from '@mui/icons-material';
import { FormControl, FormControlLabel, InputLabel, MenuItem, Paper, Select, SelectChangeEvent, Switch } from '@mui/material';
import Avatar from '@mui/material/Avatar';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Grid from '@mui/material/Grid';
import TextField from '@mui/material/TextField';
import Typography from '@mui/material/Typography';
import * as React from 'react';
import DigResponse from '../models/digResponse';
import DNSRecordType from '../models/dnsRecordTypes';
import { AppService } from '../services/app.service';

export default function DigForm() {
    const [model, setModel] = React.useState({ type: 1, domain: '' });
    const [isFormValid, setIsFormValid] = React.useState(false)
    const [showRawView, setShowRawView] = React.useState(false)
    const [showResult, setShowResult] = React.useState(false)
    const [result, setResult] = React.useState<DigResponse>({ records: [], raw: "" });
    const appService = new AppService();

    const handleTypeChange = (event: SelectChangeEvent) => {
        setModel({ ...model, type: Number(event.target.value) });
    };

    const handleDomainChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setModel({ ...model, domain: event.target.value })
        setIsFormValid(model.domain.length >= 4)
    }

    React.useEffect(() => {
        setIsFormValid(model.domain.length >= 4)
    }, [model]);

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        setShowResult(true)
        const response = await appService.Dig(model)
        setResult(response)
    };

    return (
        <>
            <Grid item xs={12} md={4} lg={3}>
                <Paper
                    sx={{
                        p: 2,
                        display: 'flex',
                        flexDirection: 'column',
                        minHeight: 400,
                    }}
                >
                    <Box
                        sx={{
                            marginTop: 8,
                            display: 'flex',
                            flexDirection: 'column',
                            alignItems: 'center',
                            flexGrow: 1,
                            minHeight: '60vh',
                            overflow: 'auto',
                        }}
                    >
                        <Avatar sx={{ m: 1, bgcolor: 'secondary.main' }}>
                            <ContentPasteSearchRounded />
                        </Avatar>
                        <Typography component="h1" variant="h5">
                            DNS Lookup
                        </Typography>
                        <Box component="form" onSubmit={handleSubmit} noValidate sx={{ mt: 1 }}>
                            <TextField
                                margin="normal"
                                required
                                fullWidth
                                id="domain"
                                label="Domain"
                                name="domain"
                                autoComplete="domain"
                                autoFocus
                                onChange={handleDomainChange}
                                error={model.domain.length < 4}
                                helperText={model.domain.length < 4 ? 'required!' : ''}
                            />
                            <FormControl fullWidth>
                                <InputLabel id="type-label">Type</InputLabel>
                                <Select
                                    labelId="type-label"
                                    id="type-select"
                                    value={model.type.toString()}
                                    label="type"
                                    onChange={handleTypeChange}
                                >
                                    {Object.values(DNSRecordType)
                                        .filter(value => typeof value === 'number')
                                        .map(value => {
                                            return <MenuItem key={value} value={value}>{DNSRecordType[value as number]}</MenuItem>
                                        })}
                                </Select>
                            </FormControl>
                            <FormControlLabel
                                control={
                                    <Switch onChange={e => setShowRawView(e.target.checked)} checked={showRawView} />
                                }
                                label="Raw View"
                            />
                            <Button
                                type="submit"
                                fullWidth
                                variant="contained"
                                sx={{ mt: 3, mb: 2 }}
                                disabled={!isFormValid}
                            >
                                go!
                            </Button>
                        </Box>
                    </Box>
                </Paper>
            </Grid>
            <Grid sx={{ display: showResult ? 'block' : 'none' }} padding={"2vh 0"} xs={12} md={4} lg={3}>
                <Paper>
                    <Box>
                        {result.records !== null ? result.records.map((record, index) =>
                            <div style={{ padding: '5px', margin: '5px' }} key={index}>{record}</div>
                        ) : "no result found!"}
                    </Box>

                </Paper>
            </Grid>

            <Grid sx={{ display: showRawView ? 'block' : 'none' }} xs={12} md={4} lg={3}>
                <Paper>
                    <Box>
                        <pre style={{ backgroundColor: '#ccc' }}>
                            {JSON.stringify(result.raw, null, '\t')}
                        </pre>
                    </Box>

                </Paper>
            </Grid>
        </>
    );
}
