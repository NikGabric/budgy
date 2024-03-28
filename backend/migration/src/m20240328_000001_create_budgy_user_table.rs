use sea_orm_migration::prelude::*;

pub struct Migration;

impl MigrationName for Migration {
    fn name(&self) -> &str {
        "m20240328_000001_create_budgy_user_table"
    }
}

#[async_trait::async_trait]
impl MigrationTrait for Migration {
    async fn up(&self, manager: &SchemaManager) -> Result<(), DbErr> {
        manager
            .create_table(
                Table::create()
                    .table(BudgyUser::Table)
                    .if_not_exists()
                    .col(
                        ColumnDef::new(BudgyUser::BudgyUserId)
                            .integer()
                            .not_null()
                            .auto_increment()
                            .primary_key(),
                    )
                    .col(ColumnDef::new(BudgyUser::Username).string().not_null())
                    .col(ColumnDef::new(BudgyUser::Password).string().not_null())
                    .col(ColumnDef::new(BudgyUser::Email).string().not_null())
                    .to_owned(),
            )
            .await
    }

    async fn down(&self, manager: &SchemaManager) -> Result<(), DbErr> {
        manager
            .drop_table(Table::drop().table(BudgyUser::Table).to_owned())
            .await
    }
}

#[derive(DeriveIden)]
pub enum BudgyUser {
    Table,
    BudgyUserId,
    Username,
    Password,
    Email,
}
