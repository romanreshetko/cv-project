import { Box, Button, FormControl, IconButton, InputAdornment, InputLabel, OutlinedInput, Stack, TextField, Typography } from "@mui/material";
import React, { useState } from "react";
import Visibility from '@mui/icons-material/Visibility';
import VisibilityOff from '@mui/icons-material/VisibilityOff';
import { useNavigate } from "react-router";

const LogIn: React.FC = () => {
    const [showPassword, setShowPassword] = useState(false);
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [errorMessage, setErrorMessage] = useState('');
    const navigate = useNavigate();

    const handleClickShowPassword = () => setShowPassword(!showPassword);
    const handleUsernameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setUsername(event.target.value);
    }
    const handlePasswordChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setPassword(event.target.value);
    }
    const handleLogin = async () => {
        if (username.length < 1 || password.length < 1) {
            setErrorMessage('Username and password fields cannot be empty.');
        } else {
            setErrorMessage('');
            
            try {
                const response = await fetch("api/auth/login", {
                    method: "POST",
                    headers: {"Content-Type": "application/json"},
                    body: JSON.stringify({username, password}),
                });

                const data = await response.json();

                if (response.ok) {
                    console.log(data.token)
                    localStorage.setItem("token", data.token)
                    navigate('/upload')
                } else {
                    setErrorMessage(data.message || "Login failed")
                }
            } catch {
                setErrorMessage("Something went wrong. Please try again.")
            }
        }
    }

    return (
        <Box sx={{
            width: '100%', height: '100vh', display: 'flex', justifyContent: 'center', alignItems: 'center',
            padding: 2, flexDirection: 'column', backgroundColor: 'linear-gradient(135deg, #f9f9f9, #eaeaea)'
            }}>
            <Typography variant='h2' gutterBottom sx={{fontWeight: 'bold', textAlign: 'center', color: '#333'}}>
                RESUME AS CODE
            </Typography>
            <Stack spacing={3} direction="column">
                <Typography variant='h6' gutterBottom sx={{textAlign: 'center', color: '#333'}}>
                    Enter Username and Password
                </Typography>
                <TextField id="username" label="Username" variant="outlined" onChange={handleUsernameChange}/>
                <FormControl sx={{m: 1}} variant='outlined'>
                    <InputLabel htmlFor='outlined-adornment-password'>Password</InputLabel>
                    <OutlinedInput
                        id="outlined-adornment-password"
                        type={showPassword ? 'text' : 'password'}
                        endAdornment={
                            <InputAdornment position='end'>
                                <IconButton
                                    aria-label={showPassword ? 'hide the password' : 'display the password'}
                                    onClick={handleClickShowPassword}
                                    edge='end'
                                >
                                    {showPassword ? <VisibilityOff /> : <Visibility />}
                                </IconButton>
                            </InputAdornment>
                        }
                        label='Password'
                        onChange={handlePasswordChange}
                    />
                </FormControl>
                {errorMessage && (
                    <Typography variant='body2' sx={{color: 'red', textAlign: 'center'}}>
                        {errorMessage}
                    </Typography>
                )}
                <Button variant="contained" sx={{fontSize: '24px', borderRadius: '20px', backgroundColor: '#4caf50',
                        color: '#fff'}} onClick={handleLogin}>
                    Log In
                </Button>
            </Stack>
        </Box>
    )
}

export default LogIn