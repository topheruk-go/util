import { isFile } from "./duck";
import { ftob } from "./ftoa";

type FormRecord = Record<string, unknown>;

export const parseForm = async (form: HTMLFormElement, map: FormRecord = {}): Promise<FormRecord> => {
    for (const [field, value] of new FormData(form)) {
        if (isFile(value)) {
            console.log(value);
            map[field] = (await ftob(value as File)).split(",")[1];
        } else {
            map[field] = value;
        }
    }
    return map;
};