<svelte:options tag="signup-form" />

<script lang="ts">
    import { parseForm } from "./form";

    export let api: string;
    export let action: string;
    export let method: string = "post";
    export let enctype: string = "multipart/form-data";

    async function submit(this: HTMLFormElement, e: Event) {
        let fr = await parseForm(this);
        fr["age"] = parseInt(fr["age"] as string);

        await fetch(api, {
            method,
            headers: [["Content_Type", "application/json"]],
            body: JSON.stringify(fr),
        });

        location.reload();
    }
</script>

<form {action} {method} {enctype} on:submit|preventDefault={submit}>
    <label>
        <p>name</p>
        <input type="text" name="name" required />
    </label>

    <label>
        <p>age</p>
        <input type="number" name="age" required />
    </label>

    <label>
        <p>file</p>
        <input type="file" name="file" accept=".pdf" required />
    </label>

    <button type="submit">submit</button>
</form>
