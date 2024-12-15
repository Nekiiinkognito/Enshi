import { EyeClosedIcon, EyeOpenIcon } from "@radix-ui/react-icons";
import { IconButton, Tooltip } from "@radix-ui/themes";
import { Dispatch, SetStateAction } from "react";

type TShowPasswordButton = {
    isShown: boolean;
    setIsShown: Dispatch<SetStateAction<boolean>>;
};

export default function ShowPasswordButton({ isShown, setIsShown }: TShowPasswordButton) {
    return (
        <div>
            <Tooltip content="Show password">
                {isShown ? (
                    <IconButton
                        type="button"
                        onClick={() => {
                            setIsShown(!isShown);
                        }}
                        size={"1"}
                        className="rounded-full"
                        variant="soft"
                    >
                        <EyeClosedIcon />
                    </IconButton>
                ) : (
                    <IconButton
                        type="button"
                        onClick={() => setIsShown(!isShown)}
                        size={"1"}
                        className="rounded-full"
                        variant="soft"
                    >
                        <EyeOpenIcon />
                    </IconButton>
                )}
            </Tooltip>
        </div>
    );
}