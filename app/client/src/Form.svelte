<svelte:options tag="signup-form" />

<script lang="ts">
    import { formattedDate } from "./form";

    //

    let fs;
    let name: string;
    let age: number;
    let email: string;
    let file: File;
    let start = formattedDate(new Date());
    let end = start;

    async function submit(this: HTMLFormElement, e: Event) {
        // let body: Record<string, FormDataEntryValue> = {};
        // for (const [k, v] of new FormData(this)) {
        //     body[k] = v;
        // }

        file = new FormData(this).get("file") as File;
        console.log(await file.arrayBuffer());
        let json = JSON.stringify({
            name,
            age,
            email,
            file: file,
        });

        await fetch("http://localhost:8000/api/v1/user", {
            method: "POST",
            headers: [["Content_Type", "application/json"]],
            body: json,
        });
    }
</script>

<form enctype="multipart/form-data" on:submit|preventDefault={submit}>
    <label>
        <p>file</p>
        <input type="file" name="file" bind:files={fs} accept=".pdf" />
    </label>

    <!-- <label>
        <p>name</p>
        <input type="text" name="name" bind:value={name} />
    </label>

    <label>
        <p>email</p>
        <input type="email" name="email" bind:value={email} />
    </label>

    <label>
        <p>age</p>
        <input type="number" name="age" bind:value={age} min="0" max="10" />
    </label> -->

    <!-- <label>
        <p>start date</p>
        <input type="date" name="start_date" bind:value={start} />
    </label>

    <label>
        <p>end date</p>
        <input type="date" name="end_date" bind:value={end} />
    </label> -->

    <button type="submit">Please save me...</button>
</form>
