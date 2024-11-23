import { createRoot } from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import './locale/i18n.ts'
import i18n from './locale/i18n.ts'

i18n.changeLanguage("ru")

createRoot(document.getElementById('root')!).render(
    <App />
)
