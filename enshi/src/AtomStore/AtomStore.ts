import { atom } from "jotai";
import { atomWithStorage } from "jotai/utils";
import { TUser } from "../@types/UserType";

export const userAtom = atom<TUser>();

export const postCreationAtom = atom<string>();
export const postCreationTitleAtom = atom<string>();

type TPostData = {
    title: string;
    content: string;
};

export const storagePostAtom = atomWithStorage<TPostData>(
    "draft-post",
    { title: "", content: "" },
    {
        getItem: (key) => sessionStorage.getItem(key) as any,
        setItem: (key, value) => sessionStorage.setItem(key, value as any),
        removeItem: (key) => sessionStorage.removeItem(key),
    },
    { getOnInit: true }
);

export const toastAtom = atom<TExistingToast[]>([]);
export const setToastAtom = atom(null, (get, set, value: TToast) => {
    let maxToastId = Math.max(...get(toastAtom).map((toast) => toast.id));
    maxToastId = maxToastId >= 0 ? maxToastId : 1;
    let atomValueWithNewToast = get(toastAtom);
    atomValueWithNewToast = [
        ...atomValueWithNewToast,
        {
            id: maxToastId + 1,
            resetFunc: (_) => {
                let currentToasts = get(toastAtom);
                let afterRemoval = currentToasts.filter(
                    (toast) => toast.id != maxToastId + 1
                );
                set(toastAtom, afterRemoval);
            },
            title: value.title,
            action: value.action,
            description: value.description,
            open: true,
        },
    ];

    set(toastAtom, atomValueWithNewToast);
});
