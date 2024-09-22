import Quill, { Delta,  } from "quill/core";
import "react-quill/dist/quill.snow.css";
import "react-quill/dist/quill.core.css";
import ReactQuill from "react-quill";
import React, {
    forwardRef,
    useEffect,
    useLayoutEffect,
    useRef,
    useState,
} from "react";
import Sources from "quill";

type TEditor = {
    readOnly?: boolean;
    defaultValue?: string | Delta;
    onTextChange?: any; // TODO: make type
    onSelectionChange?: any; // TODO same as before
};

const modules = {
    toolbar: [
        [{ header: [1, 2, 3, false] }],
        ["bold", "italic", "underline", "strike", "blockquote"],
        [
            { list: "ordered" },
            { list: "bullet" },
            { indent: "-1" },
            { indent: "+1" },
        ],
        ["link", "image"],
        ["clean"],
        [{ align: [] }],
    ],
};

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

    const changeHandler = (value: string, delta: Delta, source: Sources, editor: ReactQuill.UnprivilegedEditor) => {
        console.log(delta);
        console.log(JSON.stringify(quill?.getContents().ops, null, 2))
        
    }

    return (
        <div className="text-editor">
            <ReactQuill
                onChange={changeHandler}
                value={value}
                ref={editor}
                theme="snow"
                placeholder="Type your thoughts here..."
                modules={modules}
            />
        </div>
    );
});

export default Editor;
