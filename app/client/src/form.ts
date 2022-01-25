export const formattedDate = (d: Date): string => `${d.getFullYear()}-${(d.getMonth() + 1)
    .toString()
    .padStart(2, "0")}-${d.getDate().toString().padStart(2, "0")}`;

export interface Loan {
    name: string,
    age: number,
    "start-date": Date;
    "end-date": Date;
}

export class Loan {
    constructor(name: string, age: string, start: string, end: string) { }
}
