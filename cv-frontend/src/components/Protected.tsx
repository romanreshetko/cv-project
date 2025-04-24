import { ReactNode } from "react";
import { useNavigate } from "react-router"

interface ProtectedProps {
    children: ReactNode;
}

const Protected: React.FC<ProtectedProps> = ({children}) => {

    const token = localStorage.getItem("token");
    const isTokenValid = token && !isTokenExpired(token);
    const navigate = useNavigate();

    if (isTokenValid) {
        return children
    } else {
        navigate('/log-in')
        return null
    }
}

const isTokenExpired = (token: string) => {
    try {
        const payload = JSON.parse(atob(token.split(".")[1]))
        if (!payload.exp) {return true}
        return Date.now() >= payload.exp * 1000;
    } catch {
        return true
    }
}

export default Protected