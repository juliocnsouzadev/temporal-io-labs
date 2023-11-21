
import {log} from "@temporalio/workflow";
export class Mapped {

    words: string[];
    constructor(words:string[]) {
        this.words = words
    }

}

export async function map(text: string): Promise<Mapped> {
    text = text.toLowerCase();
    let words = text.split(/\W+/);
    return new Mapped(words);
}