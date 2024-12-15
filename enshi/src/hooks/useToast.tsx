import { useSetAtom } from "jotai";
import { setToastAtom } from "../AtomStore/AtomStore";

export default function useToast() {
    const createToast = useSetAtom(setToastAtom);
    return createToast;
}
