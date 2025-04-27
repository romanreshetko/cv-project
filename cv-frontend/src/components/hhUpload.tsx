import { useEffect, useState } from "react"

const HhUpload = () => {
    const [result, setResult] = useState('');
    useEffect(() => {
        const queryParams = new URLSearchParams(window.location.search);
        const code = queryParams.get('code');

        if (code) {
            setResult(code);
        } else {
            setResult('No code');
        }
    }, []);

    return (
        <div>
            ${result}
        </div>
    )
}

export default HhUpload