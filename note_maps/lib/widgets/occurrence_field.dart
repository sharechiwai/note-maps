// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:provider/provider.dart';

import '../controllers/controllers.dart';
import 'future_text_field.dart';

export 'future_text.dart';
export 'future_text_field.dart';

class OccurrenceField extends StatelessWidget {
  final bool autofocus;

  OccurrenceField({this.autofocus = false});

  @override
  Widget build(BuildContext context) {
    var controller = Provider.of<OccurrenceController>(context);
    return ValueListenableBuilder(
      valueListenable: controller,
      builder: (context, occurrenceState, _) => ListTile(
        title: FutureTextField(
          controller.valueTextController,
          textCapitalization: TextCapitalization.sentences,
          style: Theme.of(context).textTheme.body2,
          autofocus: autofocus,
        ),
      ),
    );
  }
}
