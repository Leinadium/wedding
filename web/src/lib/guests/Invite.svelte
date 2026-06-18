<script lang="ts">
  import { fade, fly } from "svelte/transition";
  import { onMount } from "svelte";
  import { api, type InviteResponse } from "../api";
  import Attendee from "./Attendee.svelte";
  import { loadStoredInvite, saveStoredInvite } from "./state";

  let {
    closeCb,
  }: {
    closeCb: () => void;
  } = $props();

  let inviteCode: string = $state("");
  let isLoading: boolean = $state(false);

  let invite: InviteResponse | undefined = $state(undefined);
  let currentNote: string = $state("");

  let buttonSave = $state("");
  function save() {
    buttonSave = "saved";
    setTimeout(() => {
      buttonSave = "";
    }, 3000);
  }

  onMount(() => {
    let invite = loadStoredInvite();
    if (invite) {
      inviteCode = invite;
    }
  });

  $effect(() => {
    if (inviteCode.length != 6) {
      invite = undefined;
      return;
    }
    isLoading = true;
    api
      .getInvite(inviteCode)
      .then((data) => {
        invite = data;
        currentNote = invite.note;
        saveStoredInvite(inviteCode);
        isLoading = false;
      })
      .catch((e) => {
        invite = undefined;
        isLoading = false;
      });
  });

  function updateAttendeeStatusFactory(i: number) {
    return (status: boolean | null) => {
      invite!.attendees[i].confirmed = status;
    };
  }

  function updateAttendeeIsChildFactory(i: number) {
    return (isChild: boolean) => {
      invite!.attendees[i].isChild = isChild;
    };
  }

  async function saveInvite() {
    // if note is different, save it
    if (currentNote !== invite?.note) {
      invite!.note = currentNote;
      await api.saveInviteNote(inviteCode, currentNote);
    }
    // for each attendee, save
    for (let i = 0; i < invite!.attendees.length; i++) {
      const attendee = invite!.attendees[i];
      await api.saveAttendee(attendee);
      save();
    }
  }
</script>

<div class="invite-wrapper" transition:fade={{ duration: 300 }}>
  <div class="invite" transition:fly={{ duration: 300, y: +150 }}>
    <span class="title great-cursive">RSVP</span>
    <div class="input cursive">
      <span>Input the code received:</span>
      <input type="text" placeholder="ABC123" bind:value={inviteCode} />
    </div>
    {#if isLoading}
      <span class="cursive">Loading...</span>
    {/if}
    {#if invite}
      <span
        class="description-content cursive"
        transition:fly={{ duration: 300, y: +100 }}
      >
        You can confirm your presence by clicking on the current status of each
        attendee.
      </span>
      <div class="content" transition:fly={{ duration: 300, y: +100 }}>
        {#each invite.attendees as attendee, i}
          <Attendee
            {attendee}
            updateStatus={updateAttendeeStatusFactory(i)}
            updateIsChild={updateAttendeeIsChildFactory(i)}
          />
        {/each}
      </div>

      <textarea
        class="note"
        transition:fly={{ duration: 300, y: +100 }}
        placeholder="Any observations or comments"
        bind:value={currentNote}
      ></textarea>
      <div
        class="confirm {buttonSave}"
        transition:fly={{ duration: 300, y: +100 }}
      >
        <input type="submit" value="Save" onclick={saveInvite} />
      </div>
    {/if}
    <button class="close" onclick={closeCb}>X</button>
  </div>
  <!-- <img src="src/assets/invite/invite.png" alt="invite" /> -->
</div>

<style>
  .invite-wrapper {
    position: fixed;
    top: 0;
    left: 0;

    width: 100vw;
    height: 100vh;

    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    align-items: center;

    background-color: rgba(0, 0, 0, 0.6);
    z-index: 999;
  }

  .invite {
    position: relative;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 1.5rem;
    padding: 2rem;
    border: 1px solid #e5e7eb;
    border-radius: 8px;
    font-family:
      system-ui,
      -apple-system,
      sans-serif;

    background-image: url("src/assets/invite/texture.png");
    background-color: #f0f0f0;
    box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.3);

    width: 80%;
    max-width: 500px;
    min-height: 500px;
    height: auto;
    max-height: 90vh;
    overflow-y: auto;
  }

  .input {
    display: flex;
    flex-direction: row nowrap;
    gap: 1rem;

    align-items: center;
    justify-content: center;
  }

  .input span {
    font-size: 1.5rem;
    color: #4b5563;
    font-weight: 300;
  }

  .input input {
    width: 80%;
    max-width: 80px;
    padding: 0.5rem;
    border: 0;
    border-bottom: 1px solid black;
    background: transparent;
    font-size: 1rem;
    color: #4b5563;
    font-weight: 300;
  }

  .description-content {
    font-size: 1.3rem;
    color: #4b5563;
    font-weight: 300;
  }

  .content {
    width: 100%;
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
    padding: 1rem 0;
    border-top: 1px solid #e5e7eb;
    border-bottom: 1px solid #e5e7eb;
  }

  .note {
    font-size: 1rem;
    color: #4b5563;
    font-weight: 300;

    width: 90%;
    height: 100px;

    text-align: left;
    vertical-align: top;
    font-family: cursive;
  }

  .confirm {
    display: flex;
    justify-content: flex-end; /* Aligns the save button to the right */
  }

  .confirm input[type="submit"] {
    padding: 0.75rem 2rem;
    background-color: #111827;
    color: white;
    border: none;
    border-radius: 6px;
    font-size: 1rem;
    font-weight: 600;
    cursor: pointer;
    transition: background-color 0.2s;
  }

  .confirm input[type="submit"]:hover {
    background-color: #374151;
  }

  .close {
    position: absolute;
    top: 1rem;
    right: 1rem;

    background: none;
    border: none;
    font-size: 1.5rem;
    color: #777;

    cursor: pointer;
  }

  .title {
    font-size: 2rem;
  }

  .great-cursive {
    font-family: "Alex Brush", cursive;
  }
  .cursive {
    font-family: "Great Vibes", cursive;
  }
</style>
