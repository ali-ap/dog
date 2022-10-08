import { Grid, Paper } from "@mui/material"

const HomeContent = (): JSX.Element => {
    return (
        <Grid container spacing={3}>
            <Grid item xs={12}>
                <Paper sx={{ p: 2, display: 'flex', flexDirection: 'column' }}>
                    <h1>Welcome</h1>
                    <h3>Dog is a simple domain information groper utility</h3>
                </Paper>
            </Grid>
        </Grid>
    )
};

export default HomeContent;