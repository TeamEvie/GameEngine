import org.jetbrains.kotlin.gradle.tasks.KotlinCompile

plugins {
    kotlin("jvm") version "1.6.10"
    application
}

group = "pw.evie.gameengine"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}

dependencies {
    implementation("dev.kord:kord-core:0.8.0-M15")
    implementation("org.slf4j:slf4j-nop:2.0.0-alpha6")
    implementation("redis.clients:jedis:4.2.3")
}

tasks.withType<KotlinCompile> {
    kotlinOptions.jvmTarget = "16"
}

application {
    mainClass.set("MainKt")
}