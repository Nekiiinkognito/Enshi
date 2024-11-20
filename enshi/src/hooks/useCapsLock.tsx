import { useEffect, useState } from "react";

export default function UseCapsLock() {

    const [isCapsLockOn, setIsCapsLockOn] = useState(false);

    useEffect(() => {
        const f = (e: KeyboardEvent) => {
            if (e.getModifierState("CapsLock")) {
                setIsCapsLockOn(true);
            } else {
                setIsCapsLockOn(false);
            }
        };

        document.addEventListener("keydown", f);

        return () => {
            document.removeEventListener("keydown", f);
        };
    }, []);


    return {
        isCapsLockOn
    }
}