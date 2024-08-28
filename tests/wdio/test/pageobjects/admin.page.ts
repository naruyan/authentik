import Page from "../pageobjects/page.js";
import { browser } from "@wdio/globals";

const CLICK_TIME_DELAY = 250;

export default class AdminPage extends Page {
    public async pageHeader() {
        return await $("ak-page-header").$('slot[name="header"]');
    }

    async openApplicationsListPage() {
        await this.open("if/admin/#/core/applications");
    }
}
