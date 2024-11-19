import { atom } from "jotai";
import { TUser } from "../types/UserType";

export const userAtom = atom<TUser>()