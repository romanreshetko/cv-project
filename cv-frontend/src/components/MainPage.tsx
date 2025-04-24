import { Box, Button, Stack, Typography } from "@mui/material"
import { useNavigate } from "react-router"

const MainPage: React.FC = () => {

    const navigate = useNavigate();

    return (
        <Box sx={{
            width: '100%', height: '100vh', display: 'flex', justifyContent: 'center', alignItems: 'center',
            padding: 2, flexDirection: 'column', backgroundColor: 'linear-gradient(135deg, #f9f9f9, #eaeaea)'
            }}>
            <Typography variant='h2' gutterBottom sx={{fontWeight: 'bold', textAlign: 'center', color: '#333'}}>
                RESUME AS CODE
            </Typography>
            <Stack spacing={3} direction="column">
                <Button variant="contained" sx={{fontSize: '24px', borderRadius: '20px', backgroundColor: '#4caf50',
                        color: '#fff'}} onClick={() => navigate('/sign-up')}>
                    Sign up
                </Button>
                <Button variant="contained" sx={{fontSize: '24px', borderRadius: '20px', backgroundColor: '#4caf50',
                        color: '#fff'}} onClick={() => navigate('/log-in')}>
                    Log in
                </Button>
            </Stack>
        </Box>
    )
}

export default MainPage