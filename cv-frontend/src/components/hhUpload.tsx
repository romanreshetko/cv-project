import { useEffect, useState } from "react"

const HhUpload = () => {
    const [token, setToken] = useState('');
    const [message, setMessage] = useState('');
    useEffect(async () => {
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

    return (
        <div>
            {token}
        </div>
    )
}

export default HhUpload