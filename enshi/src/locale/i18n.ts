import i18n from "i18next";
import { initReactI18next } from "react-i18next";
import ru from "./ru";
import en from "./en";

i18n
    .use(initReactI18next)
    .init({
        fallbackLng: "ru",
        lng: "ru",
        debug: true,

        resources: {
            ru: {
                translation: {...ru},
            },

            en: {
                translation: {...en},
            },
        },

        interpolation: {
            escapeValue: false,
        },
    });

export default i18n;
