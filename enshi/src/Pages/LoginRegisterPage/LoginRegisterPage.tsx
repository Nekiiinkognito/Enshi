import { useState } from "react";
import LoginElement from "./LoginElement/LoginElement";

export default function LoginRegisterPage() {
    const [isRegister, setIsRegister] = useState(false)

    return (
        <LoginElement  />
    )
}