<script lang="ts">
  import { fade, fly } from "svelte/transition";
  import { onMount } from "svelte";
  import { api, type InviteResponse } from "../api";
  import { loadStoredInvite, saveStoredInvite } from "./state";
  import Text from "../text/Text.svelte";
  import { getText } from "../text/text";
  import Submit from "./Submit.svelte";
  import Attendee from "./Attendee.svelte";

  let {
    closeCb,
  }: {
    closeCb: () => void;
  } = $props();

  let inviteCode: string = $state("");
  let isLoading: boolean = $state(false);

  let invite: InviteResponse | undefined = $state(undefined);
  let currentNote: string = $state("");

  let isSuccess: boolean = $state(false);

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

  function updateAttendee(i: number, status: boolean | null) {
    invite!.attendees[i].confirmed = status;
    console.log(i, invite!.attendees[i].confirmed);
  }

  async function saveInvite() {
    // if note is different, save it
    if (currentNote !== invite?.note) {
      invite!.note = currentNote;
      await api.saveInviteNote(inviteCode, currentNote);
    }
    // for each attendee, save
    for (let i = 0; i < invite!.attendees.length; i++) {
      var attendee = invite!.attendees[i];
      // hardcodding as adult
      attendee.isChild = false;
      await api.saveAttendee(attendee);

      isSuccess = true;
      setTimeout(() => {
        isSuccess = false;
      }, 3000);
    }
  }
</script>

<div class="invite-wrapper" transition:fade={{ duration: 300 }}>
  <div class="invite" transition:fly={{ duration: 300, y: +150 }}>
    <span class="title cursive"><Text key="invite-title" /></span>
    <div class="input formal">
      <span><Text key="invite-input" /></span>
      <input
        class="input-code formal-num"
        type="text"
        placeholder="ABC123"
        bind:value={inviteCode}
      />
    </div>
    {#if isLoading}
      <span class="formal"><Text key="invite-loading" /></span>
    {/if}
    {#if invite}
      <span
        class="description-content formal"
        transition:fly={{ duration: 300, y: +100 }}
      >
        <Text key="invite-description" />
      </span>
      <div class="content" transition:fly={{ duration: 300, y: +100 }}>
        {#each invite.attendees as attendee, i (attendee.id)}
          <Attendee index={i} {attendee} updateStatus={updateAttendee} />
        {/each}
      </div>

      <textarea
        class="note"
        transition:fly={{ duration: 300, y: +100 }}
        placeholder={getText("invite-comments")}
        bind:value={currentNote}
      ></textarea>
      <div class="confirm" transition:fly={{ duration: 300, y: +100 }}>
        <Submit success={isSuccess} onClick={saveInvite} />
      </div>
    {/if}
    <button class="close" onclick={closeCb}>&times;</button>
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

    /*background-color: rgba(0, 0, 0, 0ß.6);*/
    background-image:
      url("data:image/svg+xml,%3Csvg viewBox='0 0 200 200' xmlns='http://www.w3.org/2000/svg'%3E%3Cfilter id='noiseFilter'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.75' numOctaves='3' stitchTiles='stitch'/%3E%3C/filter%3E%3Crect width='100%25' height='100%25' filter='url(%23noiseFilter)' opacity='0.08'/%3E%3C/svg%3E"),
      radial-gradient(circle at 50% 0%, #c79eadd0 0%, #a67b8bd0 100%);
    z-index: 999;
  }
  .invite {
    position: relative;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 1.5rem;
    padding: 2rem;
    border: 2px solid #dedacd;
    border-radius: 8px;

    background-image: url("../../assets/invite/background.png");
    background-color: #a67b8bff;
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
    font-size: 1.2rem;
    color: #dedacd;
    font-weight: 300;
  }

  .input-code {
    width: 80%;
    max-width: 80px;
    padding: 0.5rem;
    border: 0;
    border-bottom: 1px solid #dedacd;
    background: transparent;
    font-size: 1rem;
    color: #dedacd;
    font-weight: 300;
  }

  .description-content {
    font-size: 1rem;
    font-style: italic;
    text-align: center;
    color: #dedacd;
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
    width: 90%;
    height: 100px;

    box-sizing: border-box;
    background-color: #00000040;
    border-radius: 1rem;
    border: 1px solid #dedacd;
    padding: 0.5rem;

    font-size: 0.9rem;
    color: #dedacd;
    font-weight: 300;
    text-align: left;
    vertical-align: top;
  }

  .confirm {
    display: flex;
    justify-content: flex-end;
  }

  .close {
    position: absolute;
    top: 1rem;
    right: 1rem;

    background: none;
    border: none;
    font-size: 1.5rem;
    color: #dedacd;

    cursor: pointer;
  }

  .title {
    color: #dedacd;
    font-size: 2rem;
  }
</style>
