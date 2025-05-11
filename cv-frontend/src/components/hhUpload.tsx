import { Box, Button, Stack, TextField, Typography } from "@mui/material";
import React, { useEffect, useState } from "react";
import { useAtom } from 'jotai';
import { fileAtom } from "../atoms/hh";

const HhUpload = () => {
    const [token, setToken] = useState('');
    const [message, setMessage] = useState('');
    const [id, setId] = useState('');
    const [file] = useAtom(fileAtom);
    useEffect(() => {
        const queryParams = new URLSearchParams(window.location.search);
        const code = queryParams.get('code');

        if (code) {
            fetch('api/hh/exchange-token', {
                method: 'POST',
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify({code})
            })
            .then(response => {
                if (!response.ok) {
                    setMessage('Failed to exchange code for token')
                    throw new Error(message);
                }
                return response.json();
            })
            .then(data => {
                setToken(data.access_token);
                setMessage('');
            })
            .catch(error => {
                console.log('Error: ', error);
            })
        } else {
            
        }
    }, []);

    const handleIdChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setId(event.target.value);
    }

    const handleUpdate = async () => {
        const formdata = new FormData();
        formdata.append("file", file);
        formdata.append("token", token);
        formdata.append("resume_id", id);

        try {
            const response = await fetch('/api/hh/update', {
                method: "PUT",
                body: formdata,
            });
    
            if (!response.ok) {
                setMessage('Failed to update resume');
                throw new Error('Failed to update resume');
            }
    
            const data = await response.json();
            console.log('Обновление успешно:', data);
            setMessage('SUCCESS!');
        } catch (error) {
            console.error('Error:', error);
            setMessage('Failed to update resume');
        }
    }

    return (
        <Box sx={{
            width: '100%', height: '100vh', display: 'flex', justifyContent: 'center', alignItems: 'center',
            padding: 2, flexDirection: 'column', backgroundColor: 'linear-gradient(135deg, #f9f9f9, #eaeaea)'
            }}>
                <Stack spacing={3} direction="column">
                    <Typography variant='h6' gutterBottom sx={{textAlign: 'center', color: '#333'}}>
                        Enter Resume Id on hh.ru
                    </Typography>
                    <TextField id="username" label="Username" variant="outlined" onChange={handleIdChange}/>
                    <Button variant="contained" sx={{fontSize: '24px', borderRadius: '20px', backgroundColor: '#4caf50',
                                                    color: '#fff'}} onClick={handleUpdate}>
                        Update resume on hh.ru
                    </Button>
                    {message && (
                        <Typography variant='body2' sx={{color: 'red', textAlign: 'center'}}>
                            {message}
                        </Typography>
                    )}
                </Stack>
        </Box>
    )
}

export default HhUpload