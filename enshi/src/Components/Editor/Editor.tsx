import Sources from "quill";
import Quill, { Delta } from "quill/core";
import {
    forwardRef,
    useEffect,
    useRef,
    useState
} from "react";
import ReactQuill from "react-quill";

type TEditor = {
    readOnly?: boolean;
    defaultValue?: string | Delta;
    onChange?: (d: string) => void;
    onSelectionChange?: any;
};

const modules = {
    toolbar: [
        [{ header: [1, 2, 3, 4, 5, false] }],
        ["bold", "italic", "underline", "strike", "blockquote", "span-wrapper"],
        [
            { list: "ordered" },
            { list: "bullet" },
            { indent: "-1" },
            { indent: "+1" },
        ],
        ["link", "image"],
        ["clean"],
        [{ align: [] }],
    ]
};

/**
 * @param onChange - function that accepts Delta element
 */
const Editor = forwardRef((props: TEditor) => {
    const editor = useRef(null);
    const [quill, setQuill] = useState<Quill | null>(null);
    const [value, setValue] = useState(new Delta())

    useEffect(() => {
        if (editor.current) {
            //@ts-ignore
            const temp = editor.current.getEditor() as Quill;
            setQuill(temp);
        }
        return () => {
            setQuill(null);
        };
    }, [editor.current]);
    

    const changeHandler = (val: string, _changeDelta: Delta, _source: Sources, _editor: ReactQuill.UnprivilegedEditor) => {
        console.log(val);
        console.log(JSON.stringify(quill?.getContents().ops, null, 2))
        let fullDelta = quill?.getContents()
       if (props.onChange) props.onChange(val || "")
        setValue(fullDelta || new Delta())
    }

    return (
        <div className="text-editor">
            <ReactQuill
                value={value}
                ref={editor}
                modules={modules}

                onChange={changeHandler}
                
                
                theme="snow"
                placeholder="Type your thoughts here..."
            />
        </div>
    );
});

export default Editor;
