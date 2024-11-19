import { MagnifyingGlassIcon } from "@radix-ui/react-icons";
import { TextField } from "@radix-ui/themes";
import { useTranslation } from "react-i18next";

export default function SearchField() {
    const {t} = useTranslation()

    return (
        <div className="flex-1">
            <TextField.Root
                className="w-full rounded-lg"
                placeholder={t("search")}
            >
                <TextField.Slot>
                    <MagnifyingGlassIcon />
                </TextField.Slot>
            </TextField.Root>
        </div>
    );
}
