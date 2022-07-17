package modules.impl

import dev.kord.core.Kord
import dev.kord.core.behavior.interaction.respondEphemeral
import dev.kord.core.entity.interaction.GuildChatInputCommandInteraction
import modules.Module

class KordPing(client: Kord) : Module(client, "KordPing", arrayOf("kord")) {
    override suspend fun onCommand(req: GuildChatInputCommandInteraction) {
        req.respondEphemeral { content = "pong!" }
    }
}
