; Make use of @local

(translation_unit (declaration (init_declarator declarator: (_) @descriptor.term)))

(namespace_definition name: (_) @descriptor.type body: (_) @descriptor.scope)
(class_specifier name: (_) @descriptor.type body: (_) @descriptor.scope)

(field_declaration declarator: (_) @descriptor.term)
(function_definition (function_declarator declarator: (_) @descriptor.method))
