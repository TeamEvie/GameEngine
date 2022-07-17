package modules.impl

import client.GameEngine
import modules.Module
import org.javacord.api.entity.message.MessageFlag
import org.javacord.api.interaction.SlashCommandInteraction

class KordPing(client: GameEngine) : Module(client, "KordPing", arrayOf("kord")) {
    override fun onCommand(req: SlashCommandInteraction) {
        req.createImmediateResponder().setContent("pong from Javacord!").setFlags(MessageFlag.EPHEMERAL).respond()
    }
}
