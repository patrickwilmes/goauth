/*
 * Copyright (c) 2024, Patrick Wilmes <p.wilmes89@gmail.com>
 * All rights reserved.
 *
 * SPDX-License-Identifier: BSD-2-Clause
 */

#include <stdio.h>
#include <sqlite3.h>
#include <stdlib.h>

void list_tables(sqlite3 *db) {
    sqlite3_stmt *stmt;
    const char *sql = "SELECT name FROM sqlite_master WHERE type='table';";
    if (sqlite3_prepare_v2(db, sql, -1, &stmt, NULL) != SQLITE_OK) {
        fprintf(stderr, "SQL error: %s\n", sqlite3_errmsg(db));
        return;
    }
    while (sqlite3_step(stmt) == SQLITE_ROW) {
        printf("%s\n", sqlite3_column_text(stmt, 0));
    }
    sqlite3_finalize(stmt);
}

void show_table_content(sqlite3 *db, const char *table_name) {
    sqlite3_stmt *stmt;
    char *sql = sqlite3_mprintf("SELECT * FROM %s;", table_name);
    if (sqlite3_prepare_v2(db, sql, -1, &stmt, NULL) != SQLITE_OK) {
        fprintf(stderr, "SQL error: %s\n", sqlite3_errmsg(db));
        sqlite3_free(sql);
        return;
    }
    while (sqlite3_step(stmt) == SQLITE_ROW) {
        for (int i = 0; i < sqlite3_column_count(stmt); i++) {
            printf("%s ", sqlite3_column_text(stmt, i));
        }
        printf("\n");
    }
    sqlite3_finalize(stmt);
    sqlite3_free(sql);
}

int main(int argc, char *argv[]) {
    sqlite3 *db;
    char *err_msg = 0;

    if (argc < 2) {
        fprintf(stderr, "Usage: %s <database>\n", argv[0]);
        return 1;
    }

    if (sqlite3_open(argv[1], &db) != SQLITE_OK) {
        fprintf(stderr, "Can't open database: %s\n", sqlite3_errmsg(db));
        sqlite3_close(db);
        return 1;
    }

    if (argc == 2) {
        list_tables(db);
    } else if (argc == 3) {
        show_table_content(db, argv[2]);
    } else {
        fprintf(stderr, "Invalid arguments\n");
    }

    sqlite3_close(db);
    return 0;
}
